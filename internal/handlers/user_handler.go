package handlers

import (
	"log"

	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"github.com/diazharizky/go-graphql-bootstrap/internal/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Age       int32
}

func (h handler) Users() []resolvers.UserResolver {
	return resolvers.NewUserList(h.appCtx)
}

func (h handler) CreateUser(args struct{ Input createUserInput }) *resolvers.UserResolver {
	user := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: args.Input.FirstName,
		LastName:  args.Input.LastName,
		Email:     args.Input.Email,
		Age:       args.Input.Age,
	}

	if err := h.appCtx.UserRepository.Create(user); err != nil {
		log.Printf("Error unable to create user record: %s", err.Error())
		return &resolvers.UserResolver{}
	}

	userResolver := resolvers.UserResolver{
		User: user,
	}

	return &userResolver
}

func (h handler) User(args struct{ ID string }) *resolvers.UserResolver {
	return resolvers.NewUser(h.appCtx, args.ID)
}
