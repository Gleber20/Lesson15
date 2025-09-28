package controller

import (
	"Lesson15/internal/config"
	"Lesson15/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
)

func (ctrl *EmployeeController) checkUserAuthentication(c *gin.Context) {

	token, err := ctrl.extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	cfg := config.LoadConfig()
	userID, isRefresh, err := pkg.ParseToken(token, cfg.AuthConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	c.Set(userIDCtx, userID)
}
