package main

import (
	"github.com/SarveshLimaye/go-rest-api/controllers"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/SarveshLimaye/go-rest-api/database"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}else{
		log.Println("Connected to MongoDB")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v1 := app.Group("/api/v1")

	v1.Get("/books", controllers.GetBooks)

	v1.Post("/books", controllers.CreateBook)

	v1.Get("/books/:id", controllers.GetBookById)

	v1.Put("/books/:id", controllers.UpdateBook)

	v1.Delete("/books/:id", controllers.DeleteBook)

	app.Listen(":3000")
}