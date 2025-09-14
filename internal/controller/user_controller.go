package controller

import (
	"Lesson15/internal/models"
	"Lesson15/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (ctrl *UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user})
}

func (ctrl *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := ctrl.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func (ctrl *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
