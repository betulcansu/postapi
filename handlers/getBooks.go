package handlers

import (
	"postapi/data"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(data.Books)
}
