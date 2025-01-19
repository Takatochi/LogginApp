package mockdata

import (
	"LoggingApp/internal/models"
	"LoggingApp/pkg/database"
	"gorm.io/driver/sqlite"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type MockDatabase struct {
	db *gorm.DB
}

// NewMockDatabase створює мок-реалізацію бази даних
func NewMockDatabase() (database.Database, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return &MockDatabase{db: db}, nil
}

// GetDB повертає екземпляр *gorm.DB
func (m *MockDatabase) GetDB() *gorm.DB {
	return m.db
}
