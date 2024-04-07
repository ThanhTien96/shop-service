package model

import "fmt"

type Currency struct {
	Currency string  `bson:"currency" validate:"require" json:"currency,omitempty"`
	Price    float32 `bson:"price" validate:"require" json:"price,omitempty"`
}

type Item struct {
	BaseModel `bson:",inline"`
	Name      string     `bson:"name" json:"name,omitempty"`
	Currency  []Currency `bson:"curency" json:"currency,omitempty"`
	Stock     int64      `bson:"stock" json:"stock,omitempty"`
}

type CreateItemRequest struct {
	Name     string     `bson:"name" validate:"require" json:"name,omitempty"`
	Currency []Currency `bson:"curency" validate:"require" json:"currency,omitempty"`
	Stock    int64      `bson:"stock" validate:"require" json:"stock,omitempty"`
}

type UpdateItemRequest struct {
	Name     string     `bson:"name" json:"name,omitempty"`
	Currency []Currency `bson:"curency" json:"currency,omitempty"`
	Stock    int64      `bson:"stock" json:"stock,omitempty"`
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
