package test

import (
	"database/sql"
	"testing"
)

func TestMockDatabase(t *testing.T) {
	// Відкриваємо базу даних у пам'яті через standard library
	sqlDB, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer sqlDB.Close()

	if err != nil {
		t.Fatalf("Failed to initialize GORM with mock database: %v", err)
	}

	// Тестова перевірка
	t.Log("Mock database initialized successfully with modernc.org/sqlite")
}
