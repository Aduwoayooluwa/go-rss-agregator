package db

import (
	"context"

	"github.com/aduwoayooluwa/go-rss-scraper/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserById(ctx context.Context, userId string) (*models.User, error) {
	collection := GetMongoClient().Database("RSS-aggr").Collection("users")

	var user models.User

	err := collection.FindOne(ctx, bson.M{"ID": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func CreateUser(ctx context.Context, user models.User) error {
	collection := GetMongoClient().Database("RSS-aggr").Collection("users")

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	collection := GetMongoClient().Database("RSS-aggr").Collection("users")

	cursor, err := collection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
