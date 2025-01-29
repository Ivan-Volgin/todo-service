package repositories

import(
	"todo-app/internal/entities"
)

type TaskRepository interface{
	Create(task entities.Task) (entities.Task, error)
	GetByID(id uint64) (entities.Task, error)
	Update(task entities.Task) (entities.Task, error)
	GetAll() ([]entities.Task, error)
	Delete(id uint64) error
}