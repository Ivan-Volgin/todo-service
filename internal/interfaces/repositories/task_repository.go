package repositories

import(
	"todo-app/internal/entities"
)

type TaskRepository interface{
	Create(task entities.Task) (entities.Task, error)
	GetByUUID(uuid string) (entities.Task, error)
	Update(task entities.Task) (entities.Task, error)
	GetAll() ([]entities.Task, error)
	Delete(uuid string) error
}