package handlers

import (
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteWorkspaceTodos(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceID")

	err := models.DeleteWorkspaceTodos(workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
