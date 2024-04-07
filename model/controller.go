package model

import (
	"shop-test/cmd/config"
	"shop-test/model"
	"shop-test/pkg/log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Controller interface {
	Load() error
	Start()
	Stop()
	Validate(i interface{}) error
	Config() *config.Config
	Logger() log.ILogger

	NewPagingRequest(page, pageSize int) *PagingRequest
	PurchaseItem(db *mongo.Database, input []model.ItemRequest, wallet []model.ResponseData) ([]model.ResponseData, error)
}
