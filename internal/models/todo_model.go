package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"description" bson:"description"`
	Owner       UserResolver       `json:"owner" bson:"owner"`
}

type TodoResolver struct {
	Todo Todo
}

func (r TodoResolver) ID() string {
	return r.Todo.ID.String()
}

func (r TodoResolver) Description() string {
	return r.Todo.Description
}

func (r TodoResolver) Owner() UserResolver {
	return r.Todo.Owner
}
