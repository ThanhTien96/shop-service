package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	C1 float32 `bson:"c1" json:"c1,omitempty"`
	C2 float32 `bson:"c2" json:"c2,omitempty"`
}

type UserItem struct {
	ItemId string `bson:"item_id" json:"item_id"`
	Count  float32  `bson:"count" json:"count"`
}

type UserResponse struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	CreatedAt time.Time         `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time         `bson:"updated_at" json:"updatedAt"`
	UserName  string             `bson:"username" json:"username,omitempty"`
	Wallets   Wallet             `json:"wallets,omitempty"`
	Items     []UserItem         `bson:"items,omitempty" json:"items"`
}

type WalletCurrency struct {
	Currency string
	Amount   float32
}
