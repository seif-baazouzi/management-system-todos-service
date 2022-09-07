package handlers

import "github.com/gofiber/fiber/v2"

func DeleteTodo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "success"})
}
