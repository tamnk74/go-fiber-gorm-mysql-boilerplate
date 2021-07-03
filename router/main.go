package router

import (
	"github.com/gofiber/fiber/v2"
	authModule "github.com/tamnk74/todolist-mysql-go/modules/auth"
	itemModule "github.com/tamnk74/todolist-mysql-go/modules/items"
)

func Init(app *fiber.App) {
	authModule.RegisterRouters(app)
	itemModule.RegisterRouters(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
