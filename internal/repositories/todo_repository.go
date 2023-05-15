package repositories

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todoRepository {
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
