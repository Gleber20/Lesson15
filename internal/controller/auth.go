package controller

import (
	"Lesson15/internal/config"
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"Lesson15/pkg"
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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (ctrl *EmployeeController) SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	accessToken, refreshToken, err := ctrl.service.Authenticate(c, models.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

const (
	refreshTokenHeader = "X-Refresh-Token"
)

func (ctrl *EmployeeController) RefreshTokenPair(c *gin.Context) {

	token, err := ctrl.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	cfg := config.LoadConfig()
	userID, isRefresh, err := pkg.ParseToken(token, cfg.AuthConfig.JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	accessToken, err := pkg.GenerateToken(
		userID,
		cfg.AuthConfig.AccessTokenTTLMinutes,
		cfg.AuthConfig.JWTSecret,
		false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: errs.ErrSomethingWentWrong.Error()})
		return
	}

	refreshToken, err := pkg.GenerateToken(
		userID,
		cfg.AuthConfig.RefreshTokenTTLDays,
		cfg.AuthConfig.JWTSecret,
		true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: errs.ErrSomethingWentWrong.Error()})
		return
	}
	
	c.JSON(http.StatusOK, SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
