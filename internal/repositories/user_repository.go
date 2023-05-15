package repositories

import "github.com/diazharizky/go-graphql-bootstrap/internal/models"

type userRepository struct{}

func NewUserRepository() userRepository {
	return userRepository{}
}

func (userRepository) List() ([]models.User, error) {
	return []models.User{}, nil
}

func (userRepository) Get(id int32) ([]models.User, error) {
	return []models.User{}, nil
}

func (userRepository) Create(params models.User) error {
	return nil
}

func (userRepository) Update(params models.User) error {
	return nil
}

func (userRepository) Delete(id int32) error {
	return nil
}
