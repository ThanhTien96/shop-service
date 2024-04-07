package main

import (
	"log"
	"shop-test/api"
	"shop-test/cmd/config"
	"shop-test/cmd/init2"
	"shop-test/controller"
	"shop-test/pkg/db"
)

func main() {
	// Parse flags
	arg := init2.LoadCommandArgs()

	// Load config
	config, err := config.LoadConfigFromFile(arg.ConfigFile)
	
	if err != nil {
		log.Panic("Failed to load config .env file", err)
	}

	// Create logger
	logger := init2.NewLogger(init2.DEV_LOG_FILE, arg.Debug, config.Logger.Encoding, "Shop-test")

	// connect db
	db.InitDB(config, logger)


	// Create controller
	ctrl := controller.NewController(config, logger, arg.Debug)

	// Create api server
	apiServer, err := controller.NewApiServer(config.Listener, logger)
	if err != nil {
		log.Fatalln(err)
	}


	api := api.NewApi(ctrl)

	// Run server
	api.Run(apiServer, logger)

}
