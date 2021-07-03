package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/repository"
)

func RegisterRouters(app fiber.Router) {
	authRepo := repository.NewAuthRepository()
	authService := NewAuthService(authRepo)
	authController := NewAutController(authService)

	router := app.Group("/api", logger.New())
	router.Post("/login", authController.login)
	router.Post("/register", authController.register)
	router.Get("/me", middlewares.AuthorizeJWT, authController.getUser)
}
