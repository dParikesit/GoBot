package main

import (
	"github.com/dParikesit/dimsBot/controllers"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("The World!")
	})

	v1 := app.Group("/api/line")
	v1.Post("/", controllers.Halo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
