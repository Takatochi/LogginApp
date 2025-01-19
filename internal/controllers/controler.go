package controllers

import (
	"LoggingApp/internal/services"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}

// BaseController — базовий контролер
type BaseController struct {
	service     services.Service
	controllers []Controller
}

// NewBaseController створює базовий контролер
func NewBaseController(service services.Service) *BaseController {
	return &BaseController{
		service:     service,
		controllers: []Controller{},
	}
}
func (b *BaseController) AddSingleController(controllerFunc func(base *BaseController) Controller) {
	b.controllers = append(b.controllers, controllerFunc(b))
}

func (b *BaseController) AddMultipleControllers(controllerFuncs ...func(base *BaseController) Controller) {
	for _, controllerFunc := range controllerFuncs {
		b.controllers = append(b.controllers, controllerFunc(b))
	}
}

func (b *BaseController) RegisterRoutes(router *gin.Engine) {
	for _, controller := range b.controllers {
		controller.RegisterRoutes(router)
	}
}
