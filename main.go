package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Addbook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{Title: "LOTR", Author: "J.R.R TOLİEN"},
	{Title: "Fahrenheit451", Author: "Ray Bradbury"},
	{Title: "Dune", Author: "Frank Herbert"},
}

func main() {

	app := fiber.New()

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	app.Post("/books", handlerAddbook)

	app.Listen(":5033")
}

func handlerAddbook(c *fiber.Ctx) error {
	addbook := &Addbook{}
	err := c.BodyParser(addbook)
	if err != nil {
		fmt.Printf("bodyparse error: %v\n", err)
		return c.SendString("error")
	}
	books = append(books, Book{Title: addbook.Title, Author: addbook.Author})
	return c.JSON(addbook)
}
