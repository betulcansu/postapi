package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "LOTR", Author: "J.R.R TOLİEN"},
	{ID: 2, Title: "Fahrenheit451", Author: "Ray Bradbury"},
	{ID: 3, Title: "Dune", Author: "Frank Herbert"},
}

func main() {

	app := fiber.New()

	app.Get("/books", handlerGetBooks)
	app.Post("/books", handlerAddbook)
	app.Delete("/books/:id", handlerDeleteBooks)

	app.Listen(":5533")
}

func handlerGetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func handlerAddbook(c *fiber.Ctx) error {
	addbook := &Book{}
	err := c.BodyParser(addbook)
	if err != nil {
		fmt.Printf("bodyparse error: %v\n", err)
		return c.SendString("error")
	}

	books = append(books, *addbook)
	return c.JSON(addbook)
}

func handlerDeleteBooks(c *fiber.Ctx) error {
	id := c.Params("id")
	iId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id.",
		})
	}
	for i, book := range books {
		if book.ID == iId {
			books = append(books[:i], books[i+1:]...)
			return c.Status(fiber.StatusOK).SendString("kitap silindi!")
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})

}
