package handlers

import (
	"fmt"
	"todos-service/src/checks"
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

	// check body
	errors := checks.CheckTodo(body)

	if errors != nil {
		return c.JSON(errors)
	}

	// create todo
	todoID, err := models.CreateTodo(body, workspaceID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "todoID": todoID})
}
