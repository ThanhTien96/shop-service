package controller

import (
	"context"
	"fmt"
	"shop-test/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Controller) PurchaseItem(db *mongo.Database, input []model.ItemRequest, wallet []model.ResponseData) ([]model.ResponseData, error) {

	shopCollection := db.Collection("shop")
	var itemIds []primitive.ObjectID
	var totalSpend []model.ResponseData
	itemCount := map[primitive.ObjectID]int32{}
	for _, item := range input {
		objID, _ := primitive.ObjectIDFromHex(item.Id)
		itemIds = append(itemIds, objID)
		itemCount[objID] = item.Count
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": itemIds, // Array of desired _id values
		},
	}

	var walletMap = map[string]float32{}
	for _, wal := range wallet {
		walletMap[wal.Currency] = wal.Amount
	}

	cursor, err := shopCollection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	var results []model.Item


	for cursor.Next(context.Background()) {
		var result model.Item
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}

		if result.Stock < itemCount[result.ID] {
			return totalSpend, fmt.Errorf(`Out of stock %d`, result.Stock)
		}

		for _, currency := range result.Currency {
			totalPrice := currency.Price * float32(itemCount[result.ID])

			if totalPrice > walletMap[currency.Currency] {
				return totalSpend, fmt.Errorf(`Out of money %s`, currency.Currency)
			}
			totalSpend = append(totalSpend, model.ResponseData{Currency: currency.Currency, Amount: totalPrice})
		}
		result.Stock -= itemCount[result.ID]
		result.UpdatedAt = time.Now()
		results = append(results, result)
	}

	for _, r := range results {
		_, err = shopCollection.UpdateOne(context.TODO(), bson.M{"_id": r.ID}, bson.D{{"$set", r}})
		if err != nil {

			return totalSpend, err
		}
	}

	return totalSpend, nil
}
