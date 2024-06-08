package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConfig struct {
	URI string
}

type MongoDBClient struct {
	Client *mongo.Client
}

func NewMongoDBClient(ctx context.Context, config MongoDBConfig) (*MongoDBClient, error) {
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("failed to connect to MongoDB %v", err)

	}

	// checking the connection
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("failed to ping mongoDB %v", err)
	}

	return &MongoDBClient{Client: client}, nil
}

func (m *MongoDBClient) Disconnect(ctx context.Context) error {
	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB: %v", err)

	}
	return nil
}
