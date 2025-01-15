package controllers

import (
	"LoggingApp/internal/models"
	"LoggingApp/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LogController залежить від logService
type LogController struct {
	logService services.LogService
	logEntry   models.LogEntry
}

// NewLogController створює новий інстанс UserController
func NewLogController(logService services.LogService) *LogController {
	return &LogController{logService: logService}
}

func (l *LogController) GetLog(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&l.logEntry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	l.logService.Logs(&l.logEntry)
	ctx.Status(http.StatusOK)
}
