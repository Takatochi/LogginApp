package migrattion

import (
	"LoggingApp/internal/models"
	Migrate2 "LoggingApp/pkg/base_migrate"
	"log"
)

type UserMigration struct {
	Migrate Migrate2.Migrate
}

func (u UserMigration) Up() error {
	log.Println("Running UserMigration...")
	return u.Migrate.Migrate(&models.User{})
}
