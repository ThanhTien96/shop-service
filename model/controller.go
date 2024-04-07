package model

import (
	"net/http"
	"shop-test/cmd/config"
	"shop-test/pkg/log"
)

type Controller interface {
	Load() error
	Start()
	Stop()
	Validate(i interface{}) error
	Config() *config.Config
	Logger() log.ILogger
	StructToMap(obj interface{}) map[string]float32

	NewPagingRequest(page, pageSize int) *PagingRequest
	ServiceGetUser(client *http.Client, token string) (*UserResponse, error) 
	PurchaseItem(input []*ItemRequest, users *UserResponse) ([]ResponseData, error)
}
