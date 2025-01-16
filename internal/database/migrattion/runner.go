package migrattion

import (
	"LoggingApp/pkg/base_migrate"
	"LoggingApp/pkg/database"
	"log"
)

func RunMigrations(store database.Database) {
	if store == nil {
		log.Fatal("Store cannot be nil")
		return
	}
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
