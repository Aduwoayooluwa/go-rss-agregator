package db

import (
	"context"
	"fmt"

	"github.com/aduwoayooluwa/go-rss-scraper/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserById(ctx context.Context, userId string) (*models.User, error) {
	collection := GetMongoClient().Database("RSS-aggr").Collection("users")

	var user models.User

	objID, _err := primitive.ObjectIDFromHex(userId)

	if _err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", _err)
	}

	_err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)

	if _err != nil {
		if _err == mongo.ErrNoDocuments {
			// Return nil to indicate no user was found, along with a custom error
			return nil, fmt.Errorf("no user found with ID %s", userId)
		}

		return nil, fmt.Errorf("error retrieving user: %w", _err)
	}

	//err := collection.FindOne(context.TODO(), bson.M{"ID": userId}).Decode(&user)

	// if err := collection.FindOne(context.TODO(), bson.M{"email": userId}).Decode(&user); err != nil {
	// fmt.Println("Error : " + err.Error())

	// }

	fmt.Println(&user)

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
