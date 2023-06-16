package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"description" bson:"description"`
	Owner       User               `json:"owner" bson:"owner"`
}
