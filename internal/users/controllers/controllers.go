package controllers

import (
	"libarymvc/internal/users/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useService models.UseService
}

func NewUserController(userService models.UseService) *UserController {
	return &UserController{useService: userService}
}

func (c *UserController) RegisterRouter(r *gin.Engine) {
	users := r.Group("/users")

	{
		users.POST("", c.CreateUser)
		users.GET("/:id", c.GetUser)
		users.GET("", c.GetAllUsers)
		users.PUT("/:id", c.UpdateUser)
		users.DELETE("/:id", c.DeleteUser)

	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.useService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	ctx.JSON(http.StatusCreated, user)

}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.useService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (c *UserController) GetAllUsers(ctx *gin.Context) {

	users, err := c.useService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.useService.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.useService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, id)

	ctx.JSON(http.StatusOK, nil)

}
