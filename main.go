package main

import (
	"fmt"

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

	app.Listen(":5033")
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
