package model

import (
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

	NewPagingRequest(page, pageSize int) *PagingRequest
}
