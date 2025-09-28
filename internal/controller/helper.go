package controller

import (
	"Lesson15/internal/config"
	"Lesson15/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func (ctrl *EmployeeController) extractTokenFromHeader(c *gin.Context, headerKey string) (string, error) {
	header := c.GetHeader(headerKey)

	if header == "" {
		return "", errors.New("empty authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty token")
	}
	return headerParts[1], nil
}

func (ctrl *EmployeeController) generateNewTokenPair(userID int) (string, string, error) {
	// генерация токенов
	cfg := config.LoadConfig()
	accessToken, err := pkg.GenerateToken(
		userID,
		cfg.AuthConfig.AccessTokenTTLMinutes,
		false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateToken(
		userID,
		cfg.AuthConfig.RefreshTokenTTLDays,
		true)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
