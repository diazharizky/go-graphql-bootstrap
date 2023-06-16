package resolvers

import "github.com/diazharizky/go-graphql-bootstrap/internal/models"

type UserResolver struct {
	User models.User
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
