package entities

import(
	"github.com/google/uuid"
)

type Task struct {
	UUID        uuid.UUID
	Title       string
	Description string
	Completed   bool
	Date		string
	Owner_ID	uint64
}
