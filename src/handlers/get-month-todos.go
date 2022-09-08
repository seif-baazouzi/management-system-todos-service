package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetMonthTodos(c *fiber.Ctx) error {
	// userID := fmt.Sprintf("%s", c.Locals("uuid"))

	// todos, err := models.GetTodayTodos(userID)

	// if err != nil {
	// 	return utils.ServerError(c, err)
	// }

	// return c.JSON(fiber.Map{"todos": todos})
	return c.JSON(fiber.Map{"todos": "list"})
}
