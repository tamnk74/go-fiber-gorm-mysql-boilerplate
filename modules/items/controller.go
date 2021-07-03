package items

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/dto"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/models"
)

type ItemController struct {
	itemService IItemService
}

type IItemController interface {
	listItems(ctx *fiber.Ctx) error
	createItem(ctx *fiber.Ctx) error
}

func (a *ItemController) listItems(c *fiber.Ctx) error {
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

func (a *ItemController) createItem(c *fiber.Ctx) error {
	var form CreateItemSchema
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
