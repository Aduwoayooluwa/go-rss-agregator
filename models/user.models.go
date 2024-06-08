package models

type User struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	Age       int    `bson:"age"`
}