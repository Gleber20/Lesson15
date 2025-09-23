package controller

import (
	"Lesson15/internal/contracts"
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	router  *gin.Engine
	service contracts.ServiceI
}

func NewUserController(service contracts.ServiceI) *EmployeeController {
	return &EmployeeController{
		router:  gin.Default(),
		service: service,
	}
}

func (ctrl *EmployeeController) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword) || errors.Is(err, errs.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}

// CreateEmployee
// @Summary Добавление нового сотрудника
// @Description Функция добавления нового сотрудника
// @Tags Employees
// @Consume json
// @Produce json
// @Param request body models.Employee true "Информация о новом сотруднике"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees [post]
func (ctrl *EmployeeController) Create(c *gin.Context) {
	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidFieldValue, err))
		return
	}
	if employee.Name == "" || employee.Email == "" || employee.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := ctrl.service.CreateEmployee(c, employee); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "Employee created successfully!"})
}

// GetEmployeeByID
// @Summary Получить сотрудника по ID
// @Description Получение конкретного сотрудника по его ID
// @Tags Employees
// @Produce json
// @Param id path int true "id сотрудника"
// @Success 200 {object} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees/{id} [get]
func (ctrl *EmployeeController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	user, err := ctrl.service.GetEmployeeByID(c, id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
}

// UpdateEmployeeByID
// @Summary Обновить сотрудника по ID
// @Description Обновление сотрудника с соответствующим ID
// @Tags Employees
// @Consume json
// @Produce json
// @Param id path int true "id сотрудника"
// @Param request body models.Employee true "Информация о сотруднике"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees/{id} [put]
func (ctrl *EmployeeController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	var employee models.Employee
	if err = c.ShouldBindJSON(&employee); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if employee.Name == "" || employee.Email == "" || employee.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	employee.ID = id

	if err = ctrl.service.UpdateEmployee(c, employee); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Employee updated successfully"})
}

// DeleteEmployeeByID
// @Summary Удалить сотрудника по ID
// @Description Удаление сотрудника по ID
// @Tags Employees
// @Produce json
// @Param id path int true "id сотрудника"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /employees{id} [delete]
func (ctrl *EmployeeController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	if err = ctrl.service.DeleteEmployee(c, id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Employee deleted successfully!"})
}
