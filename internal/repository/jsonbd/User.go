package jsonbd

import (
	"LoggingApp/internal/models"
)

type UserRepo struct {
	store *Repository
}

func (u UserRepo) CreateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetUserByID(id uint) (*models.User, error) {

	return nil, nil
}

func (u UserRepo) UpdateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}
