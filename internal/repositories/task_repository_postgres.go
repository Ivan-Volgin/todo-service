package repositories

import(
	"todo-service/internal/entities"
	"database/sql"
	"github.com/google/uuid"
	"context"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository{
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(ctx context.Context, task entities.Task) (entities.Task, error) {
	task.UUID = uuid.New().String()

	query := `INSERT INTO tasks (uuid, title, description, completed) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, task.UUID, task.Title, task.Description, task.Completed)
	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (r *PostgresTaskRepository) GetByUUID(ctx context.Context, uuid string) (entities.Task, error) {
	var task entities.Task

	query := `SELECT uuid, title, description, completed FROM tasks WHERE uuid = $1`
	row := r.db.QueryRowContext(ctx, query, uuid)

	err := row.Scan(&task.UUID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Task{}, ErrTaskNotFound
		}
		return entities.Task{}, err
	}

	return task, nil
}

func (r *PostgresTaskRepository) Update(ctx context.Context, task entities.Task) (entities.Task, error) {
	query := "UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, task.Title, task.Description, task.Completed, task.UUID)
	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (r *PostgresTaskRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Task, error) {
	var tasks []entities.Task

	query := "SELECT uuid, title, description, completed FROM tasks LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entities.Task
		err := rows.Scan(&task.UUID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

func (r *PostgresTaskRepository) Delete(ctx context.Context, uuid string) error {
	query := "DELETE FROM tasks WHERE uuid = $1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	if err != nil {
		return err
	}
	return nil
}