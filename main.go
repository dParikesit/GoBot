package main

import (
	"github.com/dParikesit/dimsBot/controllers"
	"github.com/dParikesit/dimsBot/utils"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env load failed")
	}

	app := fiber.New()
	utils.ConnectLine()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("The World!")
	})

	app.Post("/api/line", adaptor.HTTPHandlerFunc(controllers.Line))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalln("Server start error")
	}
}
