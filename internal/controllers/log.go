package controllers

import (
	"LoggingApp/internal/models"
	"LoggingApp/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LogController залежить від logService
type LogController struct {
	BaseController
	logService services.LogService
	logEntry   models.LogEntry
}

func NewLogController(base *BaseController) Controller {
	return &LogController{
		logService: base.service,
	}
}

func (l *LogController) GetLog(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&l.logEntry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	l.logService.Logs(&l.logEntry)
	ctx.Status(http.StatusOK)
}

func (l *LogController) RegisterRoutes(router *gin.Engine) {

	router.POST("/logs", l.GetLog)
}
