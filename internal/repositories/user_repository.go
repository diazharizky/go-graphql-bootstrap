package repositories

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

var users = []models.User{
	{
		ID:        1,
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foo.bar@example.com",
		Age:       15,
	},
}

func NewUserRepository(db *gorm.DB) userRepository {
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
	params.ID = int32(len(users)) + 1
	users = append(users, params)
	return nil
}

func (userRepository) Update(params models.User) error {
	return nil
}

func (userRepository) Delete(id int32) error {
	return nil
}
