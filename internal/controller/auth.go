package controller

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ctrl *EmployeeController) SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if err := ctrl.service.CreateUser(c, models.User{
		FullName: input.FullName,
		Username: input.Username,
		Password: input.Password,
	}); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully!"})
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func (ctrl *EmployeeController) SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	token, err := ctrl.service.Authenticate(c, models.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, SignInResponse{
		Token: token,
	})
}
