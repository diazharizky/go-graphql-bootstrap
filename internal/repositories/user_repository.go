package repositories

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

var users = []models.User{
	{
		ID:        primitive.NewObjectID(),
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foo.bar@example.com",
		Age:       15,
	},
}

func NewUserRepository(db *mongo.Database) userRepository {
	return userRepository{
		db: db,
	}
}

func (userRepository) List() ([]models.User, error) {
	return users, nil
}

func (userRepository) Get(id int32) ([]models.User, error) {
	return []models.User{}, nil
}

func (r userRepository) Create(params models.User) error {
	params.ID = primitive.NewObjectID()
	users = append(users, params)
	return nil
}

func (userRepository) Update(params models.User) error {
	return nil
}

func (userRepository) Delete(id int32) error {
	return nil
}
