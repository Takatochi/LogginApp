package handler

import (
	"LoggingApp/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	router *gin.Engine
	//services services.Service
}

func NewHandler(services services.Service) *Handler {
	return &Handler{
		router: gin.New(),
		//services: services,
	}
}

func (h *Handler) Routing() *gin.Engine {
	return h.router
}

//	func (h *Handler) Services() services.Service {
//		return h.services
//	}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
