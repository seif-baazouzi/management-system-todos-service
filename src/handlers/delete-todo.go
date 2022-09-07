package handlers

import (
	"fmt"
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteTodo(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	todoID := c.Params("todoID")

	// check if todo exist
	exist, err := models.IsTodoExist(todoID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !exist {
		return c.Status(404).JSON(fiber.Map{"message": "This todo does not exist"})
	}

	// delete todo
	err = models.DeleteTodo(todoID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
