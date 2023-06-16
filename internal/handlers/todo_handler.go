package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"github.com/diazharizky/go-graphql-bootstrap/internal/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler) Todos() []resolvers.TodoResolver {
	user := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: "Adi",
		LastName:  "Hidayat",
		Email:     "adi.hidayat@gmail.com",
		Age:       30,
	}

	todos := []models.Todo{
		{
			ID:          primitive.NewObjectID(),
			Description: "Todo A",
			Owner:       user,
		},
		{
			ID:          primitive.NewObjectID(),
			Description: "Todo B",
			Owner:       user,
		},
		{
			ID:          primitive.NewObjectID(),
			Description: "Todo C",
			Owner:       user,
		},
	}

	resolver := make([]resolvers.TodoResolver, len(todos))

	for i, t := range todos {
		resolver[i] = resolvers.TodoResolver{
			Todo: t,
		}
	}

	return resolver
}
