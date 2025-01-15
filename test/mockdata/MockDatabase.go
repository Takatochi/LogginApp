package mockdata

import (
	"LoggingApp/pkg/database"
	"fmt"
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
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize in-memory database: %w", err)
	}

	return &MockDatabase{db: db}, nil
}

// GetDB повертає екземпляр *gorm.DB
func (m *MockDatabase) GetDB() *gorm.DB {
	return m.db
}
