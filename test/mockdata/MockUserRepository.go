package mockdata

import (
	"LoggingApp/internal/models"
	"LoggingApp/internal/repository"
	"fmt"
)

type MockRepository struct {
	users  map[uint]models.User
	nextID uint
}

func NewMockRepository() repository.UserRepository {
	return &MockRepository{
		users:  make(map[uint]models.User),
		nextID: 1,
	}
}
func (m MockRepository) CreateUser(user *models.User) error {
	user.ID = m.nextID
	m.users[m.nextID] = *user
	m.nextID++
	return nil
}

func (m MockRepository) GetUserByID(id uint) (*models.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return &user, nil
}

func (m MockRepository) UpdateUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRepository) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}
