package test

import (
	"LoggingApp/internal/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
func TestWithMock(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %s", err)
	}
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {
			t.Fatalf("Error closing mock database: %s", err)
		}
	}(mockDB)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error initializing GORM with sqlmock: %s", err)
	}

	// Налаштовуємо очікування запитів
	mock.ExpectQuery("SELECT \\* FROM \"users\"").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Test Name"),
	)

	// Ваш код тестування тут
	var result []models.User
	db.Find(&result)

	// Перевіряємо, що всі очікування виконані
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not all expectations were met: %s", err)
	}
}
