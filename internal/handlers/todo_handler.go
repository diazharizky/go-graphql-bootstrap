package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler) Todos() []models.TodoResolver {
	user := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: "Adi",
		LastName:  "Hidayat",
		Email:     "adi.hidayat@gmail.com",
		Age:       30,
	}

	userResolver := models.UserResolver{
		User: user,
	}

	todos := []models.Todo{
		{
			ID:          primitive.NewObjectID(),
			Description: "Todo A",
			Owner:       userResolver,
		},
		{
			ID:          primitive.NewObjectID(),
			Description: "Todo B",
			Owner:       userResolver,
		},
		{
			ID:          primitive.NewObjectID(),
			Description: "Todo C",
			Owner:       userResolver,
		},
	}

	resolver := make([]models.TodoResolver, len(todos))

	for i, t := range todos {
		resolver[i] = models.TodoResolver{
			Todo: t,
		}
	}

	return resolver
}
