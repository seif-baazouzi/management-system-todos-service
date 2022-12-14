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

	app.Get("/api/v1/todos/today", auth.IsLogin, handlers.GetTodayTodos)
	app.Get("/api/v1/todos/remaining", auth.IsLogin, handlers.GetRemainingTodos)
	app.Get("/api/v1/todos/:workspaceID/month/:month", auth.IsWorkspaceOwner, handlers.GetMonthTodos)
	app.Post("/api/v1/todos/:workspaceID", auth.IsWorkspaceOwner, handlers.CreateTodo)
	app.Put("/api/v1/todos/:todoID", auth.IsLogin, handlers.UpdateTodo)
	app.Delete("/api/v1/todos/:todoID", auth.IsLogin, handlers.DeleteTodo)
	app.Delete("/api/v1/todos/workspace/:workspaceID", auth.IsWorkspaceOwner, handlers.DeleteWorkspaceTodos)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
