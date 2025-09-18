package controller

import (
	"Lesson15/internal/models"
	"Lesson15/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	router  *gin.Engine
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		router:  gin.Default(),
		service: service,
	}
}

// CreateUser
// @Summary Добавление нового пользователя
// @Description Функция добавления нового пользователя
// @Tags Users
// @Consume json
// @Produce json
// @Param request body models.User true "Информация о новом пользователе"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
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

// GetUserByID
// @Summary Получить пользователя по ID
// @Description Получение конкретного пользователя по его ID
// @Tags Users
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user})
}

// UpdateUserByID
// @Summary Обновить пользователя по ID
// @Description Обновление пользователя с соответствующим ID
// @Tags Users
// @Consume json
// @Produce json
// @Param id path int true "id пользователя"
// @Param request body models.User true "Информация о пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
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

// DeleteUserByID
// @Summary Удалить пользователя по ID
// @Description Удаление пользователя по ID
// @Tags Users
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
