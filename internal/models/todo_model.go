package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      string             `json:"userId" bson:"userId"`
	Description string             `json:"description" bson:"description"`
}
