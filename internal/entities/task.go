package entities

import (
	"github.com/google/uuid"
)

type Task struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Date        string    `json:"date"`
	User_ID     uint64    `json:"user_id"`
}
