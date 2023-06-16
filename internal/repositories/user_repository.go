package repositories

import (
	"context"

	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(db *mongo.Database) userRepository {
	return userRepository{
		coll: db.Collection("users"),
	}
}

func (repo userRepository) List() ([]models.User, error) {
	ctx := context.TODO()

	cur, err := repo.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepository) Get(id int32) ([]models.User, error) {
	return []models.User{}, nil
}

func (repo userRepository) Create(newUser models.User) error {
	_, err := repo.coll.InsertOne(context.TODO(), newUser)
	return err
}

func (userRepository) Update(params models.User) error {
	return nil
}

func (userRepository) Delete(id int32) error {
	return nil
}
