package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Age       int32              `json:"age" bson:"age"`
}
