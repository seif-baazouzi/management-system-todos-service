package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	TodoID      uuid.UUID `json:"todoID"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Done        bool      `json:"done"`
	UserID      uuid.UUID `json:"userID"`
	WorkspaceID uuid.UUID `json:"workspaceID"`
	CreatedAt   time.Time `json:"createdAt"`
}

type TodoBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}
