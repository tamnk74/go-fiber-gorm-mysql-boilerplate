package Jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tamnk74/todolist-mysql-go/config"
	"github.com/tamnk74/todolist-mysql-go/models"
)

func GenerateAccessToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Duration(config.JWT_EXP) * time.Second,
		"time":  time.Now(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		panic(err)
	}

	return tokenString
}

func VerifyAccessToken(tokenString string) (models.User, bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic("Unexpected signing method")
		}

		return []byte(config.JWT_SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return models.User{
			ID:    uint(claims["id"].(float64)),
			Email: claims["email"].(string),
		}, true
	}

	return models.User{}, false
}
