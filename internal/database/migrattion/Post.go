package migrattion

import (
	"LoggingApp/pkg/base_migrate"
	"gorm.io/gorm"
	"log"
)

// example test this register in provider migrattion
type Post struct {
	gorm.Model
	Title   string `gorm:"size:255"`
	Content string `gorm:"type:text"`
}
type PostMigration struct {
	Migrate base_migrate.Migrate
}

func (p PostMigration) Up() error {
	log.Printf("Running PostMigration...")
	return p.Migrate.Migrate(&Post{})
}
