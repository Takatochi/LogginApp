package services

import (
	"LoggingApp/internal/models"
	"log"
)

// LogService інтерфейс для управління користувачами
type LogService interface {
	Logs(logEntry *models.LogEntry)
}

// logService реалізація LogService
type logService struct {
}

// NewUserService створює новий інстанс UserService
func NewlogService() LogService {
	return &logService{}
}

func (l *logService) Logs(logEntry *models.LogEntry) {

	log.Printf("Level: %s, Message: %s,Category: %s, Time: %s\n", logEntry.Level, logEntry.Message,
		logEntry.Category, logEntry.Timestamp)

}
