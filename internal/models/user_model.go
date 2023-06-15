package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Age       int32              `json:"age" bson:"age"`
}

type UserResolver struct {
	User User
}

func (r UserResolver) ID() string {
	return r.User.ID.String()
}

func (r UserResolver) FirstName() string {
	return r.User.FirstName
}

func (r UserResolver) LastName() string {
	return r.User.LastName
}

func (r UserResolver) Email() string {
	return r.User.Email
}

func (r UserResolver) Age() int32 {
	return r.User.Age
}
