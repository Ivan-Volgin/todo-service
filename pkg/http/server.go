package http

import(
	"github.com/gorilla/mux"
	"todo-app/internal/interfaces/handlers"
	"todo-app/internal/interfaces/repositories"
	"todo-app/internal/usecases"
)

func NewServer(taskRepo repositories.TaskRepository) *mux.Router {
	r := mux.NewRouter()

	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUseCase)

	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{uuid}", taskHandler.GetByUUID).Methods("GET")
	r.HandleFunc("/tasks/{uuid}", taskHandler.Update).Methods("PUT")
	r.HandleFunc("/tasks/{uuid}", taskHandler.Delete).Methods("DELETE")
	r.HandleFunc("/tasks", taskHandler.GetAll).Methods("GET")

	return r
}