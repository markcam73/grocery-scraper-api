package controllers

import (
	"grocery-scraper-api/models"
	"grocery-scraper-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users := c.userService.GetUsers()
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser := c.userService.CreateUser(user)
	ctx.JSON(http.StatusCreated, createdUser)
}
