package handlers

import (
	"encoding/json"
	"net/http"
	"todo-service/internal/entities"
	"todo-service/internal/usecases"
	"todo-service/internal/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

// @Summary Create a task
// @Description Create a task
// @Tags tasks
// @Accept json
// @Produce json
// @Success 201 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tasks/create [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := h.taskUseCase.CreateTask(r.Context(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

// @Summary Get task by uuid
// @Description Get task by uuid
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tasks/{uuid} [get]
func (h *TaskHandler) GetByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"] // проверить что значение есть

	task, err := h.taskUseCase.GetByUUID(r.Context(), uuid)
	if err != nil{
		if err == repositories.ErrTaskNotFound {
            http.Error(w, "Task not found", http.StatusNotFound)
            return
        }
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// @Summary Get task by uuid
// @Description Get task by uuid
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tasks/{uuid}/update [patch]
func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask, err := h.taskUseCase.UpdateTask(r.Context(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

// @Summary Get all tasks
// @Description Get all user's tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array}  entities.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{user_uuid}/tasks [get]
func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, offset := 10, 0

    if limitStr != "" {
        l, err := strconv.Atoi(limitStr)
        if err != nil {
            http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
            return
        }
        limit = l
    }

    if offsetStr != "" {
        o, err := strconv.Atoi(offsetStr)
        if err != nil {
            http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
            return
        }
        offset = o
    }

	tasks, err := h.taskUseCase.GetAll(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Delete task by uuid
// @Description Delete task by its uuid
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tasks/{uuid}/delete [delete]
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if err := h.taskUseCase.DeleteTask(r.Context(), uuid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}