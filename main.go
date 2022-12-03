package main

import (
	"postapi/handlers"

	"github.com/gofiber/fiber/v2"
)




func main() {

	app := fiber.New()

	app.Get("/books", handlers.GetBooks)
	app.Post("/books", handlers.AddBook)
	app.Delete("/books/:id", handlers.DeleteBook)
	app.Put("/books/:id", handlers.UpdateBook)

	app.Listen(":5530")
}
