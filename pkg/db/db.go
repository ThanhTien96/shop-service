package db

import (
	"context"
	"fmt"
	"shop-test/cmd/config"
	"shop-test/pkg/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return DB.Collection(col)
}

// NewMongoStore postgres init
func InitDB(cfg *config.Config, logger log.ILogger) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s?authSource=admin",
		cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName,
	)
	logger.Info(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()


	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Panicf("ERR: DB Connect Error ")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Panicf("ERR: DB Connect Error ")
	}

	DB = client.Database(cfg.DB.DBName)

	logger.Info("INF: MongoDB connecte successfully!.")

}

func CloseDB(logger log.ILogger) error {
	logger.Infof("INF: Mongodb disconect")
	return DB.Client().Disconnect(context.Background())
}
