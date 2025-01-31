package usecases

import(
	"todo-service/internal/entities"
	"todo-service/internal/interfaces/repositories"
)

type TaskUseCase struct {
	repo repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) *TaskUseCase{
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) CreateTask(task entities.Task) (entities.Task, error){
	return uc.repo.Create(task)
}

func (uc *TaskUseCase) GetByUUID(uuid string) (entities.Task, error){
	task, err := uc.repo.GetByUUID(uuid)
    if err != nil {
        if err == repositories.ErrTaskNotFound {
            return entities.Task{}, err
        }
        return entities.Task{}, err
    }
    return task, nil
}

func (uc *TaskUseCase) UpdateTask(task entities.Task) (entities.Task, error){
	return uc.repo.Update(task)
}

func (uc *TaskUseCase) GetAll() ([]entities.Task, error){
	return uc.repo.GetAll()
}

func (uc *TaskUseCase) DeleteTask(uuid string) error{
	return uc.repo.Delete(uuid)
}