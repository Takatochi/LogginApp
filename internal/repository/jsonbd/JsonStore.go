package jsonbd

import "LoggingApp/internal/repository"

type Repository struct {
	userRepository *UserRepo
}

func NewRepository() *Repository {
	return &Repository{}
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
