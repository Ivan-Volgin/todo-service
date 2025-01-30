package repositories

import(
	"todo-app/internal/entities"
	"database/sql"
	"github.com/google/uuid"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository{
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(task entities.Task) (entities.Task, error) {
	task.UUID = uuid.New().String()

	query := `INSERT INTO tasks (uuid, title, description, completed) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, task.UUID, task.Title, task.Description, task.Completed)
	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (r *PostgresTaskRepository) GetByUUID(uuid string) (entities.Task, error) {
	var task entities.Task

	query := `SELECT * FROM tasks WHERE uuid = $1`
	row := r.db.QueryRow(query, uuid)

	err := row.Scan(&task.UUID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Task{}, ErrTaskNotFound
		}
		return entities.Task{}, err
	}

	return task, nil
}