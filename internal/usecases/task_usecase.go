package usecases

import (
	"errors"
	"fmt"
	"log"
	"context"
	"todo-service/internal/entities"
	"todo-service/internal/repositories"
)

type TaskUseCase struct {
	repo repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) *TaskUseCase {
	log.Println("TaskUseCase initialized")
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, task entities.Task) (entities.Task, error) {
	return uc.repo.Create(ctx, task)
}

func (uc *TaskUseCase) GetByUUID(ctx context.Context, uuid string) (entities.Task, error) {
	task, err := uc.repo.GetByUUID(ctx, uuid)
	if err != nil {
		if errors.Is(err, repositories.ErrTaskNotFound) {
			return entities.Task{}, fmt.Errorf("error task by uuid %s; %w", uuid, err)
		}
		return entities.Task{}, err
	}
	return task, nil
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, task entities.Task) (entities.Task, error) {
	return uc.repo.Update(ctx, task)
}

func (uc *TaskUseCase) GetAll(ctx context.Context, limit, offset int) ([]entities.Task, error) {
	return uc.repo.GetAll(ctx, limit, offset)
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, uuid string) error {
	return uc.repo.Delete(ctx, uuid)
}
