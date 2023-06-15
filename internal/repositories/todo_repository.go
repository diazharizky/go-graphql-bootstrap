package repositories

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	db *mongo.Database
}

func NewTodoRepository(db *mongo.Database) todoRepository {
	return todoRepository{
		db: db,
	}
}

func (todoRepository) List() ([]models.Todo, error) {
	return []models.Todo{}, nil
}

func (todoRepository) Get(id int32) ([]models.Todo, error) {
	return []models.Todo{}, nil
}

func (todoRepository) Create(params models.Todo) error {
	return nil
}

func (todoRepository) Update(params models.Todo) error {
	return nil
}

func (todoRepository) Delete(id int32) error {
	return nil
}
