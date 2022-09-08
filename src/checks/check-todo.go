package checks

import (
	"todos-service/src/models"

	"github.com/gofiber/fiber/v2"
)

func CheckTodo(body models.TodoBody) fiber.Map {
	errors := make(fiber.Map)

	if body.Title == "" {
		errors["title"] = "Must not be empty"
	}

	if body.StartingDate.IsZero() {
		errors["startingDate"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}
