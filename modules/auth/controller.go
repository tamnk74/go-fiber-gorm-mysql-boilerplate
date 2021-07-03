package auth

import (
	"github.com/gocraft/work"
	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/constants"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/utils/queue"
)

type AuthController struct {
	authService IAuthService
}

type IAuthController interface {
	login(ctx *fiber.Ctx) error
	register(ctx *fiber.Ctx) error
	getUser(ctx *fiber.Ctx) error
}

func (a *AuthController) login(c *fiber.Ctx) error {
	var form LoginSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	token, err := a.authService.Login(form.Email, form.Password)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"access_token": token,
			"token_type":   "bearer",
		},
	})
}

func (a *AuthController) getUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	err := a.authService.GetUser(&user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": user,
	})
}

func (a *AuthController) register(c *fiber.Ctx) error {
	var form RegisterSchema
	if err := c.BodyParser(&form); err != nil {
		return err
	}

	errors := middlewares.ValidateStruct(form)
	if len(errors) != 0 {
		return c.Status(int(errors[0].Status)).JSON(errors)
	}

	user := models.User{Name: form.Name, Email: form.Email, Password: form.Password}
	newItem, err := a.authService.CreateUser(user)
	if err != nil {
		return err
	}
	queue.CreateJob(constants.SEND_EMAIL_Q, work.Q{
		"subject": "Welcome " + user.Name + " to Go App",
		"email":   user.Email,
	})
	return c.Status(200).JSON(fiber.Map{"data": newItem})
}
