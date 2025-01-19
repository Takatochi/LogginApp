package services

import "LoggingApp/internal/repository"

type Service interface {
	LogService
	UserService
}

// service реалізація об'єднаного Service
type service struct {
	UserService
	LogService
}

// NewService створює новий об'єднаний Service
func NewService(repo repository.Store) Service {
	return &service{
		UserService: NewUserService(repo),
		LogService:  NewlogService(),
	}
}
