package app

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

type IUserRepository interface {
	List() ([]models.User, error)
	Get(id int) ([]models.User, error)
	Create(params models.User) error
	Update(params models.User) error
	Delete(id int) error
}
