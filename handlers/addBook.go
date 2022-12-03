package handlers

import (
	"fmt"
	"postapi/data"
	"postapi/models"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {

	addbook := &models.Book{}
	err := c.BodyParser(addbook)
	if err != nil {
		fmt.Printf("bodyparse error: %v\n", err)
		return c.SendString("error")
	}

	data.Books = append(data.Books, *addbook)

	return c.JSON(fiber.Map{
		"status": "kitap eklendi!",
		"fields": data.Books,
	})
}
