package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	router *gin.Engine
}

func NewHandler() *Handler {
	return &Handler{
		router: gin.New(),
	}
}

func (h *Handler) Routing() *gin.Engine {
	return h.router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
