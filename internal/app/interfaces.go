package app

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	List() ([]models.User, error)
	Get(id string) (*models.User, error)
	Create(params *models.User) error
	Update(params models.User) error
	Delete(id string) error
}

type ITodoRepository interface {
	List(filter bson.M) ([]models.Todo, error)
	Get(id string) (*models.Todo, error)
	Create(params *models.Todo) error
	Update(params models.Todo) error
	Delete(id string) error
}
