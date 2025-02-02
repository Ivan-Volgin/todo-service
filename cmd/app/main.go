package main

import (
	"log"
	"net/http"
	"todo-service/internal/interfaces/handlers"
	"todo-service/internal/interfaces/repositories"
	"todo-service/internal/usecases"
	"todo-service/pkg/database"
	"todo-service/pkg/http_server"
)

func main(){
	db, err := database.Conect("localhost", "5432", "postgres", "123", "todo_db")
	if err != nil {
		log.Fatalf("Could not connect to database %v", err)
	}
	defer db.Close()

	taskRepo := repositories.NewPostgresTaskRepository(db)

	taskUseCase := usecases.NewTaskUseCase(taskRepo)

	taskHandler := handlers.NewTaskHandler(taskUseCase)

	r := http_server.NewServer(*taskHandler)

	log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}