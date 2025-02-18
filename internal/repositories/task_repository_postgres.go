package repositories

import (
	"context"
	"log"
	"database/sql"
	"github.com/google/uuid"
	"todo-service/internal/entities"
)

const(
	db_name = "tasks"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	log.Println("PostgresTaskRepository initialized")
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(ctx context.Context, task entities.Task) (entities.Task, error) {
	task.UUID = uuid.New()

	query := `INSERT INTO tasks (uuid, title, description, completed, user_id, date) VALUES ($1, $2, $3, $4, $5, $6)`
	log.Printf("Creating task with UUID: %s", task.UUID)
	_, err := r.db.ExecContext(ctx, query, task.UUID, task.Title, task.Description, task.Completed, task.Date, task.User_ID)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return entities.Task{}, err
	}
	log.Printf("Task created successfully with UUID: %s", task.UUID)
	return task, nil
}

func (r *PostgresTaskRepository) GetByUUID(ctx context.Context, uuid string) (entities.Task, error) {
	var task entities.Task

	query := `SELECT uuid, title, description, completed, date, user_id FROM tasks WHERE uuid = $1`
	log.Printf("Fetching task by UUID: %s", uuid)
	row := r.db.QueryRowContext(ctx, query, uuid)

	err := row.Scan(&task.UUID, &task.Title, &task.Description, &task.Completed, &task.Date, &task.User_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Task not found for UUID: %s", uuid)
			return entities.Task{}, ErrTaskNotFound
		}
		log.Printf("Error fetching task by UUID: %v", err)
		return entities.Task{}, err
	}
	log.Printf("Task fetched successfully with UUID: %s", task.UUID)
	return task, nil
}

func (r *PostgresTaskRepository) Update(ctx context.Context, task entities.Task) (entities.Task, error) {
	query := "UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4"
	log.Printf("Updating task with UUID: %s", task.UUID)
	_, err := r.db.ExecContext(ctx, query, task.Title, task.Description, task.Completed, task.UUID)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		return entities.Task{}, err
	}
	log.Printf("Task updated successfully with UUID: %v", err)
	return task, nil
}

func (r *PostgresTaskRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Task, error) {
	var tasks []entities.Task

	query := "SELECT uuid, title, description, completed FROM tasks LIMIT $1 OFFSET $2"
	log.Printf("Fetching all tasks with limit: %d, offset: %d", limit, offset)
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entities.Task
		err := rows.Scan(&task.UUID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Printf("Error scanning task row: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over task rows: %v", err)
		return nil, err
	}

	log.Printf("Fetched %d tasks successfully", len(tasks))
	return tasks, nil
}

func (r *PostgresTaskRepository) Delete(ctx context.Context, uuid string) error {
	query := "DELETE FROM tasks WHERE uuid = $1"
	log.Printf("Deleting task with UUID: %s", uuid)
	_, err := r.db.ExecContext(ctx, query, uuid)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return err
	}
	log.Printf("Task successfully deleted with UUID: %v", uuid)
	return nil
}

// func (r *PostgresTaskRepository) GetByName(ctx context.Context, owner_id, task_title string) (entities.Task, error) {
// 	var task entities.Task

// }

// func (r *PostgresTaskRepository) GetByDate(ctx context.Context, owner_id, date string) (entities.Task, error) {

// }
