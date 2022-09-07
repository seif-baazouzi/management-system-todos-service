package main

import (
	"fmt"
	"os"
	"todos-service/src/auth"
	"todos-service/src/db"
	"todos-service/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.InitDataBase()
	defer db.CloseDataBase()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/v1/todos/:workspaceID/today", auth.IsWorkspaceOwner, handlers.GetTodayTodos)
	app.Post("/api/v1/todos/:workspaceID", auth.IsWorkspaceOwner, handlers.CreateTodo)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
