package controllers

import (
	"LoggingApp/internal/models"
	"LoggingApp/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController залежить від UserService
type UserController struct {
	userService services.UserService
}

// NewUserController створює новий інстанс UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	users := uc.userService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser := uc.userService.CreateUser(user)
	ctx.JSON(http.StatusCreated, newUser)
}
