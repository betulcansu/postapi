package handlers

import (
	"postapi/data"
	"postapi/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateBook(c *fiber.Ctx) error {

	var booksBody models.Book

	err := c.BodyParser(&booksBody)

	if err != nil {
		return c.SendString("hata: body okunamadi")
	}

	id := c.Params("id")
	intID, _ := strconv.Atoi(id)

	if data.Books[intID-1].Title != "" {
		data.Books[intID-1] = booksBody
	} else {
		return c.SendString("bu id yok!")
	}

	return c.JSON(data.Books)
}
