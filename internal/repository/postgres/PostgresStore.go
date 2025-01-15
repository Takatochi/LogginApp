package postgres

import (
	"LoggingApp/internal/repository"
	"gorm.io/gorm"
)

type Repository struct {
	db             *gorm.DB
	userRepository *UserRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) User() repository.UserRepository {
	if r.userRepository != nil {
		return r.userRepository
	}

	r.userRepository = &UserRepo{
		store: r,
	}

	return r.userRepository
}
