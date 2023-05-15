package app

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

type IUserRepository interface {
	List() ([]models.User, error)
	Get(id int32) ([]models.User, error)
	Create(params models.User) error
	Update(params models.User) error
	Delete(id int32) error
}

type ITodoRepository interface {
	List() ([]models.Todo, error)
	Get(id int32) ([]models.Todo, error)
	Create(params models.Todo) error
	Update(params models.Todo) error
	Delete(id int32) error
}
