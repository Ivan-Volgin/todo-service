package http_server

import(
	"github.com/gorilla/mux"
	"todo-service/internal/handlers"
)

func NewServer(taskHandler handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{uuid}", taskHandler.GetByUUID).Methods("GET")
	r.HandleFunc("/tasks/{uuid}", taskHandler.Update).Methods("PATCH")
	r.HandleFunc("/tasks/{uuid}", taskHandler.Delete).Methods("DELETE")
	r.HandleFunc("/tasks", taskHandler.GetAll).Methods("GET")

	return r
}