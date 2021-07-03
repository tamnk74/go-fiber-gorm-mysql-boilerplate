package items

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/repository"
)

func RegisterRouters(app fiber.Router) {
	itemRepo := repository.NewItemRepository()
	itemService := NewItemService(itemRepo)
	itemController := NewItemController(itemService)

	router := app.Group("/api").Use(middlewares.AuthorizeJWT)

	router.Get("/items", itemController.listItems)
	router.Post("/items", itemController.createItem)
}
