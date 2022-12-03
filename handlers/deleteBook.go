package handlers

import (
	"postapi/data"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteBook(c *fiber.Ctx) error {


	id := c.Params("id")
	iId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id.",
		})
	}
	for i, book := range data.Books {
		if book.ID == iId {
			data.Books = append(data.Books[:i], data.Books[i+1:]...)

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": "kitap silindi!",
				"fields": data.Books,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})

}
