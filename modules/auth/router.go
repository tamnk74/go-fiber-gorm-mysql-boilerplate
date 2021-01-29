package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
)

func RegisterRouters(app fiber.Router) {
	authController := NewAuthController()
	app.Post("/login", authController.login)
	app.Post("/register", authController.register)
	app.Get("/me", middlewares.AuthorizeJWT, authController.getUser)
}
