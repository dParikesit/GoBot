package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Line(c *fiber.Ctx) error {
	var result map[string]interface{}
	err := json.Unmarshal(c.Body(), &result)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return c.SendString("Halo")
}
