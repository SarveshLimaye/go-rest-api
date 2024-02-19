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

	app.Get("/api/v1/books", controllers.GetBooks)

	app.Post("/api/v1/books", controllers.CreateBook)

	app.Get("/api/v1/books/:id", controllers.GetBookById)

	app.Put("/api/v1/books/:id", controllers.UpdateBook)

	app.Delete("/api/v1/books/:id", controllers.DeleteBook)

	app.Listen(":3000")
}