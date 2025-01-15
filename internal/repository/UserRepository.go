package repository

import "LoggingApp/internal/models"

// UserRepository визначає інтерфейс для роботи з користувачами
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}
