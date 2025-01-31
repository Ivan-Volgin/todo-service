package repositories

import(
	"todo-service/internal/entities"
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

func (r *PostgresTaskRepository) Update(task entities.Task) (entities.Task, error) {
	query := "UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4"
	_, err := r.db.Exec(query, task.Title, task.Description, task.Completed, task.UUID)
	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (r *PostgresTaskRepository) GetAll() ([]entities.Task, error) {
	var tasks []entities.Task

	query := "SELECT * FROM tasks"
	rows, err := r.db.Query(query)
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

func (r *PostgresTaskRepository) Delete(uuid string) error {
	query := "DELETE FROM tasks WHERE uuid = $1"
	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}
	return nil
}