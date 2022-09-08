package handlers

import (
	"fmt"
	"regexp"
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetMonthTodos(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	workspaceID := c.Params("workspaceID")

	// validate month
	month := c.Params("month")

	matched, err := regexp.Match("^[0-9]{4}-[0-9]{1,2}$", []byte(month))

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !matched {
		return c.JSON(fiber.Map{"message": "invalid-month-format"})
	}

	// get todos
	todosLists := make(map[string][]models.Todo)

	for _, day := range utils.GetMonthDays(month) {
		today := fmt.Sprintf("%s-%d", month, day)

		todos, err := models.GetTodos(userID, today, workspaceID)

		if err != nil {
			return utils.ServerError(c, err)
		}

		todosLists[today] = todos
	}

	return c.JSON(fiber.Map{"todos": todosLists})
}
