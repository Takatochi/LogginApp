package test

import (
	"LoggingApp/internal/models"
	"LoggingApp/test/mockdata"
	"testing"
)

func TestCreateUser(t *testing.T) {
	repo := mockdata.NewMockRepository()
	newUser := models.User{Name: "Test User", Email: "test@example.com"}

	err := repo.CreateUser(&newUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}
