package handlers

import (
	"fmt"
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateTodo(c *fiber.Ctx) error {
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

	// parse body
	var body models.TodoBody

	err = c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check values
	if body.Title == "" {
		return c.JSON(fiber.Map{"title": "Must not be empty"})
	}

	// update todo
	err = models.UpdateTodo(body, todoID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
