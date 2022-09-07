package handlers

import (
	"fmt"
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	workspaceID := c.Params("workspaceID")

	// parse body
	var body models.TodoBody

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check values
	if body.Title == "" {
		return c.JSON(fiber.Map{"title": "Must not be empty"})
	}

	// create todo
	todoID, err := models.CreateTodo(body, workspaceID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "todoID": todoID})
}
