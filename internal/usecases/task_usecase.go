package usecases

import(
	"todo-app/internal/entities"
	"todo-app/internal/interfaces/repositories"
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

func (uc *TaskUseCase) GetBuID(id uint64) (entities.Task, error){
	return uc.repo.GetByID(id)
}

func (uc *TaskUseCase) UpdateTask(task entities.Task) (entities.Task, error){
	return uc.repo.Update(task)
}

func (uc *TaskUseCase) GetAll() ([]entities.Task, error){
	return uc.repo.GetAll()
}

func (uc *TaskUseCase) DeleteTask(id uint64) error{
	return uc.repo.Delete(id)
}