package middlewares

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tamnk74/todolist-mysql-go/dto"
)

func HandleApiError() fiber.Config {
	return fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Statuscode defaults to 500
			log.Println(err.Error())
			var apiErrors []*dto.ApiError
			apiError, ok := err.(*dto.ApiError)
			if ok {
				apiErrors = append(apiErrors, apiError)
				return ctx.Status(int(apiError.Status)).JSON(fiber.Map{
					"errors": apiErrors,
				})
			}

			// Handle fiber error
			e, ok := err.(*fiber.Error)
			if ok {
				apiErrors = append(apiErrors, &dto.ApiError{
					Status: http.StatusBadRequest,
					Code:   strconv.Itoa(e.Code),
					Title:  e.Message,
					Detail: e.Message,
				})
				return ctx.Status(e.Code).JSON(fiber.Map{
					"errors": apiErrors,
				})
			}

			if err != nil {
				apiErrors = append(apiErrors, &dto.ApiError{
					Status: http.StatusInternalServerError,
					Code:   "ERR-" + strconv.Itoa(http.StatusInternalServerError),
					Title:  err.Error(),
					Detail: err.Error(),
				})
				return ctx.Status(500).JSON(fiber.Map{
					"errors": apiErrors,
				})
			}

			// Return from handler
			return nil
		},
	}
}
