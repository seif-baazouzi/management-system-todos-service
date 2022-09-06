package handlers

import (
	"fmt"
	"time"
	"todos-service/src/models"
	"todos-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTodayTodos(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceID")

	date := time.Now()
	today := fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
	fmt.Println(today)
	todos, err := models.GetTodos(workspaceID, today, today)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"todos": todos})
}
