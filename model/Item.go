package model

import "fmt"

type Currency struct {
	Currency string  `bson:"currency" validate:"require" json:"currency,omitempty"`
	Price    float32 `bson:"price" validate:"require" json:"price,omitempty"`
}

type Item struct {
	BaseModel `bson:",inline"`
	Name      string     `bson:"name" json:"name"`
	Currency  []Currency `bson:"currency" json:"currency"`
	Stock     int64      `bson:"stock" json:"stock"`
}

type CreateItemRequest struct {
	Name     string     `bson:"name,omitempty" validate:"require" json:"name,omitempty"`
	Currency []Currency `bson:"curency,omitempty" validate:"require" json:"currency,omitempty"`
	Stock    int64      `bson:"stock,omitempty" validate:"require" json:"stock,omitempty"`
}


type ItemRequest struct {
	ItemId string `json:"item_id" validate:"required"`
	Count int64 `json:"count" validate:"required"`
}

type PurchaseItemRequest struct {
	Items []ItemRequest `json:"items" validate:"required"`
}



type ResponseData struct {
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

func (ireq *ItemRequest) ValidateRequest() error {
	if ireq.ItemId == "" {
		return fmt.Errorf("item id is required")
	}

	if ireq.Count == 0 {
		return fmt.Errorf("count is required")
	}

	return nil
}



func (i *CreateItemRequest) ValidateItem() error {
	if i.Name == "" {
		return fmt.Errorf("Item name is required")
	}

	for _, item := range i.Currency {
		if item.Currency == "" {
			return fmt.Errorf("Item Currency.currency is required")
		}

		if item.Price == 0 {
			return fmt.Errorf("Item Currency.Price is required")
		}
	}

	if i.Stock == 0 {
		return fmt.Errorf("Item stock is required")
	}

	return nil
}
