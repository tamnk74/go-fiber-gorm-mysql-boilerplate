package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	authModule "github.com/tamnk74/todolist-mysql-go/modules/auth"
	itemModule "github.com/tamnk74/todolist-mysql-go/modules/items"
)

func Init(app *fiber.App) {

	// Public api
	publicApi := app.Group("/api", logger.New())
	authModule.RegisterRouters(publicApi)

	// Authenticated api
	authApi := app.Group("/api").Use(middlewares.AuthorizeJWT)
	itemModule.RegisterRouters(authApi)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
