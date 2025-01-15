package migrattion

import (
	"LoggingApp/pkg/base_migrate"
	"LoggingApp/pkg/database"
	"log"
)

func RunMigrations(store database.Database) {
	migrator := base_migrate.NewMigrate(store)

	// Виконуємо всі міграції
	migrations := getMigrations(migrator)
	for _, migration := range migrations {
		err := migration.Up()
		if err != nil {
			log.Printf("Migration failed: %v\n", err)
			continue
		}
		log.Println("Migration succeeded!")
	}
}
