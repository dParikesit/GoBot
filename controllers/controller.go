package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dParikesit/dimsBot/utils"
	"github.com/gofiber/fiber/v2"
)

func Line(c *fiber.Ctx) error {
	var result utils.WebhookEvent
	err := json.Unmarshal(c.Body(), &result)
	if err != nil {
		return err
	}

	fmt.Println(result.Events.Type)
	return c.SendString("Halo")
}
