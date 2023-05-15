package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

type createUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Age       int32
}

func (h handler) Users() []models.UserResolver {
	users, err := h.appCtx.UserRepository.List()
	if err != nil {
		return []models.UserResolver{}
	}

	resolver := make([]models.UserResolver, len(users))

	for i, u := range users {
		resolver[i] = models.UserResolver{
			User: u,
		}
	}

	return resolver
}

func (handler) CreateUser(args struct{ Input createUserInput }) *models.UserResolver {
	user := models.User{
		FirstName: args.Input.FirstName,
		LastName:  args.Input.LastName,
		Email:     args.Input.Email,
		Age:       args.Input.Age,
	}

	userResolver := models.UserResolver{
		User: user,
	}

	return &userResolver
}
