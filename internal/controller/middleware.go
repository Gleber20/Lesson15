package controller

import (
	"Lesson15/internal/config"
	"Lesson15/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
)

func (ctrl *EmployeeController) checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty authorization header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid authorization header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty token",
		})
		return
	}

	cfg := config.LoadConfig()
	token := headerParts[1]
	userID, err := pkg.ParseToken(token, cfg.AuthConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set(userIDCtx, userID)
	c.Next()
}
