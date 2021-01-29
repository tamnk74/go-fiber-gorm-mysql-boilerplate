package items

import "github.com/gofiber/fiber/v2"

func RegisterRouters(app fiber.Router) {
	itemController := NewItemController()
	app.Get("/items", itemController.listItems)
	app.Post("/items", itemController.createItem)
}
