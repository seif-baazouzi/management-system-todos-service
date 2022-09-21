package models

import (
	"database/sql"
	"todos-service/src/db"

	"github.com/google/uuid"
)

func GetTodos(userID string, date string, workspace ...string) ([]Todo, error) {
	workspaceID := ""
	if len(workspace) > 0 {
		workspaceID = workspace[0]
	}

	conn := db.GetPool()
	defer db.ClosePool(conn)

	var todos []Todo

	var rows *sql.Rows
	var err error

	if workspaceID != "" {
		rows, err = conn.Query(
			`SELECT todoID, title, body, done, startingDate, endingDate, userID, workspaceID, createdAt FROM todos
			WHERE userID = $1 AND workspaceID = $3
			AND (endingDate = '0001-01-01' AND startingDate = $2)
			OR (endingDate <> '0001-01-01' AND startingDate <= $2 AND endingDate >= $2)
			ORDER BY createdAt`,
			userID,
			date,
			workspaceID,
		)
	} else {
		rows, err = conn.Query(
			`SELECT todoID, title, body, done, startingDate, endingDate, userID, workspaceID, createdAt FROM todos
			WHERE userID = $1
			AND (endingDate = '0001-01-01' AND startingDate = $2)
			OR (endingDate <> '0001-01-01' AND startingDate <= $2 AND endingDate >= $2)
			ORDER BY createdAt`,
			userID,
			date,
		)
	}

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.TodoID, &todo.Title, &todo.Body, &todo.Done, &todo.StartingDate, &todo.EndingDate, &todo.UserID, &todo.WorkspaceID, &todo.CreatedAt)

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
		"INSERT INTO todos (todoID, title, body, done, startingDate, endingDate, workspaceID, userID) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		todoID,
		todo.Title,
		todo.Body,
		todo.Done,
		todo.StartingDate,
		todo.EndingDate,
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
		"UPDATE todos SET title = $1, body = $2, done = $3, startingDate = $4, endingDate = $5 WHERE todoID = $6 AND userID = $7",
		todo.Title,
		todo.Body,
		todo.Done,
		todo.StartingDate,
		todo.EndingDate,
		todoID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(todoID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"DELETE FROM todos WHERE todoID = $1 AND userID = $2",
		todoID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteWorkspaceTodos(workspaceID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"DELETE FROM todos WHERE workspaceID = $1",
		workspaceID,
	)

	if err != nil {
		return err
	}

	return nil
}
