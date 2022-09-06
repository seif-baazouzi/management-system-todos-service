package models

import "todos-service/src/db"

func GetTodos(workspaceID string, startingDate string, endingDate string) ([]Todo, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var todos []Todo

	rows, err := conn.Query(
		"SELECT todoID, title , body , userID, workspace, createdAt FROM todos WHERE workspace = $1 AND createdAt between $2 AND $3",
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
		rows.Scan(&todo.TodoID, &todo.Title, &todo.Body, &todo.UserID, &todo.WorkspaceID, &todo.CreatedAt)

		todos = append(todos, todo)
	}

	return todos, nil
}
