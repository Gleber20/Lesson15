package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserID    int  `json:"user_id"`
	isRefresh bool `json:"is_refresh"`
}

func GenerateToken(userID int, ttl int, secret string, isRefresh bool) (string, error) {

	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{},
		UserID:         userID,
		isRefresh:      isRefresh,
	}
	if isRefresh {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * 24 * time.Hour)
	} else {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * time.Minute)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string, secret string) (int, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, false, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserID, claims.isRefresh, nil
	}

	return 0, false, fmt.Errorf("invalid token")
}
