package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func ConnectMongoDB(ctx context.Context, config MongoDBConfig) error {
	var err error

	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(config.URI)
		clientInstance, err = mongo.Connect(ctx, clientOptions)

		if err != nil {
			err = fmt.Errorf("failed to connect to MongoDB: %v", err)
			return
		}

		//  checking the conncection
		err = clientInstance.Ping(ctx, readpref.Primary())
		if err != nil {
			err = fmt.Errorf("failed to ping mongoDN: %v", err)
			return
		}
	})

	return err
}

// Get mongoclient returns in singleton mongoDB client instance
func GetMongoClient() *mongo.Client {
	return clientInstance
}

// disconnect mongodb and closing the db conncection
func DisconnectMongoDB(ctx context.Context) error {
	if clientInstance != nil {
		if err := clientInstance.Disconnect(ctx); err != nil {
			return fmt.Errorf("failed to disconnect MongoDB: %v", err)
		}
	}

	return nil
}
