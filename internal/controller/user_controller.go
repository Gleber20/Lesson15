package controller

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"Lesson15/internal/service"
	"errors"
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

func (ctrl *UserController) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotfound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
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
		ctrl.handleError(c, errors.Join(errs.ErrInvalidFieldValue, err))
		return
	}
	if user.Name == "" || user.Email == "" || user.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := ctrl.service.CreateUser(user); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully!"})
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		ctrl.handleError(c, err)
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	var user models.User
	if err = c.ShouldBindJSON(&user); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if user.Name == "" || user.Email == "" || user.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	user.ID = id

	if err = ctrl.service.UpdateUser(user); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Product updated successfully"})
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	if err = ctrl.service.DeleteUser(id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "User deleted successfully!"})
}
