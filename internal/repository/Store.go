package repository

// Repository implement from interface Store
type Store interface {
	User() UserRepository
}
