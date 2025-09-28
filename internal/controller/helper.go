package controller

import (
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
