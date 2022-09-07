package models

import (
	"todos-service/src/db"

	"github.com/google/uuid"
)

func GetTodos(workspaceID string, startingDate string, endingDate string) ([]Todo, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var todos []Todo

	rows, err := conn.Query(
		"SELECT todoID, title, body, done, userID, workspaceID, createdAt FROM todos WHERE workspaceID = $1 AND createdAt between $2 AND $3",
		workspaceID,
		startingDate,
		endingDate,
	)

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.TodoID, &todo.Title, &todo.Body, &todo.Done, &todo.UserID, &todo.WorkspaceID, &todo.CreatedAt)

		todos = append(todos, todo)
	}

	return todos, nil
}

func IsTodoExist(todoID string, userID string) (bool, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	rows, err := conn.Query("SELECT 1 FROM todos WHERE todoID = $1 AND userID = $2", todoID, userID)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func CreateTodo(todo TodoBody, workspaceID string, userID string) (string, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	todoID := uuid.New()

	_, err := conn.Exec(
		"INSERT INTO todos (todoID, title, body, done, workspaceID, userID) VALUES ($1, $2, $3, $4, $5, $6)",
		todoID,
		todo.Title,
		todo.Body,
		todo.Done,
		workspaceID,
		userID,
	)

	if err != nil {
		return "", err
	}

	return todoID.String(), nil
}

func UpdateTodo(todo TodoBody, todoID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"UPDATE todos SET title = $1, body = $2, done = $3 WHERE todoID = $4 AND userID = $5",
		todo.Title,
		todo.Body,
		todo.Done,
		todoID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}
