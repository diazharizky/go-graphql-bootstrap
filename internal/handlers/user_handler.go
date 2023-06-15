package handlers

import (
	"log"

	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		log.Printf("Error unable to retrieve user records: %s", err.Error())
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

func (h handler) CreateUser(args struct{ Input createUserInput }) *models.UserResolver {
	user := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: args.Input.FirstName,
		LastName:  args.Input.LastName,
		Email:     args.Input.Email,
		Age:       args.Input.Age,
	}

	if err := h.appCtx.UserRepository.Create(user); err != nil {
		log.Printf("Error unable to create user record: %s", err.Error())
		return &models.UserResolver{}
	}

	userResolver := models.UserResolver{
		User: user,
	}

	return &userResolver
}
