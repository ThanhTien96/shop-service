package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"shop-test/api"
	"shop-test/model"
	"shop-test/pkg/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) PurchaseItem(input []*model.ItemRequest, users *model.UserResponse) ([]model.ResponseData, error) {
    var totalSpend []model.ResponseData
    var resultItems []model.Item
    coll := db.GetDBCollection(api.ItemColName)

    var itemIds []primitive.ObjectID
    itemCount := map[primitive.ObjectID]int64{}

    for _, item := range input {
        objID, _ := primitive.ObjectIDFromHex(item.ItemId)
        itemIds = append(itemIds, objID)
        itemCount[objID] = item.Count
    }

    filter := bson.M{
        "_id": bson.M{
            "$in": itemIds,
        },
    }

    // user wallet
    var userWalletList = map[string]float32{}
    objValue := c.StructToMap(users.Wallets)

    for currency, amount := range objValue {
        userWalletList[currency] = amount
    }

    cursor, err := coll.Find(context.Background(), filter)
    if err != nil {
        panic(err)
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var result model.Item
        err := cursor.Decode(&result)
        if err != nil {
            panic(err)
        }

        // check quantity
        if result.Stock < itemCount[result.ID] {
            return totalSpend, fmt.Errorf(`out of stock %d`, result.Stock)
        }

        // Calculate total spend for each currency
        for _, currency := range result.Currency {
            totalPrice := currency.Price * float32(itemCount[result.ID])

            walletBalance, ok := userWalletList[currency.Currency]
            
            if !ok {
                return nil, fmt.Errorf("currency %s not found in user's wallets", currency.Currency)
            }

            if totalPrice > walletBalance {
                return nil, fmt.Errorf("insufficient balance for currency %s", currency.Currency)
            }
            totalSpend = append(totalSpend, model.ResponseData{Currency: currency.Currency, Amount: totalPrice})
        }

        // Update stock and append the updated item to resultItems
        result.UpdatedAt = time.Now()
        result.Stock -= itemCount[result.ID]
        resultItems = append(resultItems, result)

    }
    fmt.Println(totalSpend)
    
    // Update the stock in the database for each updated item
    for _, r := range resultItems {
        _, err = coll.UpdateOne(context.TODO(), bson.M{"_id": r.ID}, bson.D{{"$set", r}})
        if err != nil {
            return totalSpend, err
        }
    }

    return totalSpend, nil
}


func (c *Controller) ServiceGetUser(client *http.Client, token string) (*model.UserResponse, error) {
	url := fmt.Sprintf("%v/auth/profile", c.config.ServiceAPI.AuthService)
	var responseData *model.UserResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return responseData, err
	}

	// set header bearer token
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return responseData, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return responseData, fmt.Errorf("server trả về mã trạng thái không hợp lệ: %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&responseData); err != nil {
		return responseData, err
	}

	return responseData, nil
}
