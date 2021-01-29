package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/dto"
	Jwt "github.com/tamnk74/todolist-mysql-go/utils/jwt"
	"github.com/tamnk74/todolist-mysql-go/utils/redis"
)

func AuthorizeJWT(c *fiber.Ctx) error {
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.Get("Authorization")
	if len(authHeader) <= 7 {
		c.Status(401).JSON(fiber.Map{
			"errors": []dto.ApiError{
				{
					Code:   "ERR-401",
					Status: http.StatusUnauthorized,
					Title:  http.StatusText(http.StatusUnauthorized),
					Detail: http.StatusText(http.StatusUnauthorized),
				},
			},
		})
		return nil
	}
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]
	user, ok := Jwt.VerifyAccessToken(tokenString)
	isExistToken := redis.IsExistToken(strconv.Itoa(int(user.ID)), tokenString)

	if ok && isExistToken {
		c.Locals("user", user)
		return c.Next()
	}

	return c.Status(401).JSON(fiber.Map{
		"errors": []dto.ApiError{
			{
				Code:   "ERR-401",
				Status: http.StatusUnauthorized,
				Title:  http.StatusText(http.StatusUnauthorized),
				Detail: http.StatusText(http.StatusUnauthorized),
			},
		},
	})
}
