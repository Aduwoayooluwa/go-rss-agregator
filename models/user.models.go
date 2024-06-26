package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	Email       string             `bson:"email"`
	Age         int                `bson:"age"`
	LastUpdated time.Time          `bson:"lastUpdated"`
	CreatedAt   time.Time          `bson:"createdAt"`
}
