package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type APIUnauthorizeError struct {
	Code    int32    `json:"code" example:"401"`
	Success bool     `json:"success" example:"false"`
	Errors  []string `json:"errors" example:"unauthorized token"`
}
type APINotFoundError struct {
	Code    int32    `json:"code" example:"404"`
	Success bool     `json:"success" example:"false"`
	Errors  []string `json:"errors" example:"something not found"`
}

type APIResponseFail struct {
	APIResponseSuccess
	Errors []string `json:"errors" example:"something error"`
}

type APIResponseSuccess struct {
	Code    int32       `json:"code" example:"200"`
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

type APIResponseDeleted struct {
	Code    int32    `json:"code" example:"204"`
	Success bool     `json:"success" example:"true"`
	Message string   `json:"message" example:"deleted"`
	Data    []string `json:"data"`
}
