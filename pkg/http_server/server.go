package http_server

import(
	"github.com/gorilla/mux"
	"todo-service/internal/handlers"
	"log"
)

func NewServer(taskHandler handlers.TaskHandler) *mux.Router {
	log.Println("Initializing HTTP server and configuring routes...")
	r := mux.NewRouter()

    log.Println("Registering POST /tasks endpoint for creating tasks")
    r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

    log.Println("Registering GET /tasks/{uuid} endpoint for retrieving tasks by UUID")
    r.HandleFunc("/tasks/{uuid}", taskHandler.GetByUUID).Methods("GET")

    log.Println("Registering PATCH /tasks/{uuid} endpoint for updating tasks")
    r.HandleFunc("/tasks/{uuid}", taskHandler.Update).Methods("PATCH")

    log.Println("Registering DELETE /tasks/{uuid} endpoint for deleting tasks")
    r.HandleFunc("/tasks/{uuid}", taskHandler.Delete).Methods("DELETE")

    log.Println("Registering GET /tasks endpoint for retrieving all tasks")
    r.HandleFunc("/tasks", taskHandler.GetAll).Methods("GET")

    log.Println("HTTP server routes configured successfully")

	return r
}