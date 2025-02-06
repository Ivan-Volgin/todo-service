package main

import (
	"log"
	"net/http"
	"os"
	// "context"
	"todo-service/internal/handlers"
	"todo-service/internal/repositories"
	"todo-service/internal/usecases"
	"todo-service/pkg/database"
	"todo-service/pkg/http_server"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
)

// @title           todo-service
// @version         1.0
// @description     This is an application for creating and managing tasks

// @contact.name   Volgin Ivan
// @contact.email  volgin.i.a@yandex.ru


// @host      localhost:8080
// @BasePath  /
func main(){
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    serverAddr := os.Getenv("SERVER_ADDR")

    connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	db, err := database.Conect(connStr)
	if err != nil {
		log.Fatalf("Could not connect to database %v", err)
	}
	defer db.Close()

	taskRepo := repositories.NewPostgresTaskRepository(db)

	taskUseCase := usecases.NewTaskUseCase(taskRepo)

	taskHandler := handlers.NewTaskHandler(taskUseCase)

	r := http_server.NewServer(*taskHandler)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(serverAddr, r))
}