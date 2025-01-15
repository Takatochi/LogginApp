package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database interface {
	GetDB() *gorm.DB
}
type GormDatabase struct {
	db *gorm.DB
}

func NewGormDatabase(cfg Config) Database {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil
	}
	log.Print("Connected to PostgreSQL successfully!")
	return &GormDatabase{db: db}
}
func (g *GormDatabase) GetDB() *gorm.DB {
	return g.db
}
