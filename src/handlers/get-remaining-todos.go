package handlers

import (
	"fmt"
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetRemainingTodos(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))

	todos, err := models.GetRemainingTodos(userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"todos": todos})
}
