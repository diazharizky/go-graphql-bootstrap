package repositories

import (
	"context"

	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo todoRepository) List(filter bson.M) ([]models.Todo, error) {
	ctx := context.TODO()

	cur, err := repo.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	if err = cur.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (repo todoRepository) Get(id string) (todo *models.Todo, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	filter := bson.M{
		"_id": objectID,
	}

	err = repo.coll.FindOne(context.TODO(), filter).Decode(&todo)

	return
}

func (repo todoRepository) Create(newTodo *models.Todo) error {
	newTodo.ID = primitive.NewObjectID()

	_, err := repo.coll.InsertOne(context.TODO(), newTodo)
	return err
}

func (todoRepository) Update(params models.Todo) error {
	return nil
}

func (todoRepository) Delete(id string) error {
	return nil
}
