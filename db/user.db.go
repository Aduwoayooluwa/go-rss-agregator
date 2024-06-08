package db

import (
	"context"

	"github.com/aduwoayooluwa/go-rss-scraper/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(ctx context.Context, user models.User) error {
	collection := GetMongoClient().Database("").Collection("users")

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	collection := GetMongoClient().Database("").Collection("users")

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
