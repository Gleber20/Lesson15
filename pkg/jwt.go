package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func GenerateToken(userID int, ttl int, secret string) (string, error) {
	expiration := time.Now().Add(time.Duration(ttl) * time.Minute).Unix()

	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration,
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string, secret string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
