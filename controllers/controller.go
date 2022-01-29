package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Line(c *fiber.Ctx) error {
	log.Println(c.Body())
	return c.SendString("Halo")
}
