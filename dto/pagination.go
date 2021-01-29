package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	Limit   int `query:"limit"`
	Page    int `query:"page"`
	PerPage int
	Total   int64
	Offset  int
}

func (a *Pagination) Default() {
	a.Page = 1
	a.Limit = 15
	a.PerPage = a.Limit
	a.Offset = 0
}

func (a *Pagination) FillDefault() {
	if a.Page == 0 {
		a.Page = 1
	}
	if a.Limit == 0 {
		a.Limit = 15
	}
	a.PerPage = a.Limit
	a.Offset = (a.Page - 1) * a.Limit
	if a.Limit == -1 {
		a.Offset = -1
	}
}

func (a *Pagination) ParseParam(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return err
	}
	a.Page = page
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil {
		return err
	}
	a.Limit = limit
	if a.Limit == 0 {
		a.Limit = 15
	}
	a.PerPage = a.Limit
	a.Offset = (a.Page - 1) * a.Limit
	if a.Limit == -1 {
		a.Offset = -1
	}
	return nil
}

func (a *Pagination) Update() {
	if a.Limit == -1 {
		a.PerPage = int(a.Total)
	}
}

func (a *Pagination) GetMeta() fiber.Map {
	return fiber.Map{
		"total":    a.Total,
		"page":     a.Page,
		"per_page": a.PerPage,
	}
}
