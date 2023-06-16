package repositories

import (
	"context"

	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	coll *mongo.Collection
}

func NewTodoRepository(db *mongo.Database) todoRepository {
	return todoRepository{
		coll: db.Collection("todos"),
	}
}

func (repo todoRepository) List() ([]models.Todo, error) {
	ctx := context.TODO()

	cur, err := repo.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	if err = cur.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (todoRepository) Get(id int32) ([]models.Todo, error) {
	return []models.Todo{}, nil
}

func (repo todoRepository) Create(newTodo models.Todo) error {
	_, err := repo.coll.InsertOne(context.TODO(), newTodo)
	return err
}

func (todoRepository) Update(params models.Todo) error {
	return nil
}

func (todoRepository) Delete(id int32) error {
	return nil
}
