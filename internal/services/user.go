package services

import (
	"LoggingApp/internal/models"
	"LoggingApp/internal/repository"
)

// UserService інтерфейс для управління користувачами
type UserService interface {
	GetAllUsers() []models.User
	CreateUser(user models.User) models.User
}

// userService реалізація UserService
type userService struct {
	users []models.User
	repo  repository.Store
}

// NewUserService створює новий інстанс UserService
func NewUserService(repo repository.Store) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers() []models.User {
	user, err := s.repo.User().GetUserByID(10)
	if err != nil {
		return nil
	}
	s.users = append(s.users, *user)
	return s.users
}

func (s *userService) CreateUser(user models.User) models.User {
	user.ID = uint(len(s.users) + 1)
	s.users = append(s.users, user)
	return user
}
