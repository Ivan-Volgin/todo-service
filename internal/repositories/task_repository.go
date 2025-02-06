package repositories

import(
	"todo-service/internal/entities"
	"context"
)

type TaskRepository interface{
	Create(ctx context.Context, task entities.Task) (entities.Task, error)
	GetByUUID(ctx context.Context, uuid string) (entities.Task, error)
	Update(ctx context.Context, task entities.Task) (entities.Task, error)
	GetAll(ctx context.Context, limit, offset int) ([]entities.Task, error)
	Delete(ctx context.Context, uuid string) error
}