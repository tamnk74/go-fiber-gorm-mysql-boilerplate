package items

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/dto"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
)

type ItemController interface {
	listItems(ctx *fiber.Ctx) error
	createItem(ctx *fiber.Ctx) error
}

type itemController struct {
	itemService ItemService
}

func NewItemController() ItemController {
	itemRepo := repository.NewItemRepository()
	itemService := NewItemService(itemRepo)
	return &itemController{
		itemService: itemService,
	}

}

func (a *itemController) listItems(c *fiber.Ctx) error {
	pagi := new(dto.Pagination)
	if err := c.QueryParser(pagi); err != nil {
		return err
	}
	pagi.FillDefault()
	books, _ := a.itemService.ListItems(pagi)
	pagi.Update()
	c.Links(
		"http://api.example.com/users?page=2", "next",
		"http://api.example.com/users?page=5", "last",
	)
	return c.Status(200).JSON(fiber.Map{
		"meta": pagi.GetMeta(),
		"data": books,
	})
}

func (a *itemController) createItem(c *fiber.Ctx) error {
	var form CreateItem
	if err := c.BodyParser(&form); err != nil {
		return err
	}

	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}

	item := models.Item{Name: form.Name}
	newItem, _ := a.itemService.CreateItem(item)
	return c.Status(200).JSON(fiber.Map{"data": newItem})
}
