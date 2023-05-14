package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

func (handler) Users() []models.UserResolver {
	users := []models.User{
		{
			ID:        1,
			FirstName: "Adi",
			LastName:  "Hidayat",
			Email:     "adi.hidayat@gmail.com",
			Age:       38,
		},
		{
			ID:        2,
			FirstName: "Syafiq",
			LastName:  "Riza",
			Email:     "syafiq.riza@gmail.com",
			Age:       45,
		},
	}

	resolver := make([]models.UserResolver, len(users))

	for i, u := range users {
		resolver[i] = models.UserResolver{
			User: u,
		}
	}

	return resolver
}
