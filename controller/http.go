package controller

import (
	"net/http"
	"time"

	"shop-test/pkg/log"
)


// NewApiServer(address string, logger log.ILogger, tlsEnabled bool, serverCrt string, serverKey string)
func NewApiServer(address string, logger log.ILogger) (*http.Server, error) {

	apiServer := &http.Server{
		Addr:      address,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	logger.Info("INF: Loading API Listener on ", address)

	return apiServer, nil
}
