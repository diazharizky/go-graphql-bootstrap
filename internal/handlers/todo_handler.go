package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

func (handler) Todos() []models.TodoResolver {
	user := models.User{
		ID:        1,
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
			ID:          1,
			Description: "Todo A",
			Owner:       userResolver,
		},
		{
			ID:          2,
			Description: "Todo B",
			Owner:       userResolver,
		},
		{
			ID:          3,
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
