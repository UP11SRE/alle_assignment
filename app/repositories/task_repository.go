package repositories

import (
	"alle_assignment/app/models"
	"alle_assignment/config"
	"fmt"
	"strings"
)

func CreateTask(task *models.Task) error {
	query := `
		INSERT INTO tasks (title, description, status, user_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	row := config.DB.QueryRow(query, task.Title, task.Description, task.Status, task.UserName)
	err := row.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	return err
}

func GetAllTasks(page int) ([]models.Task, int, error) {
	limit := 10
	offset := (page - 1) * limit

	query := `SELECT id, title, description, status, user_name, created_at, updated_at
	          FROM tasks ORDER BY id ASC LIMIT $1 OFFSET $2
`

	rows, err := config.DB.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID, &task.Title, &task.Description, &task.Status,
			&task.UserName, &task.CreatedAt, &task.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, task)
	}

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM tasks`
	err = config.DB.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}



func UpdateTask(id int, updates map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	argPos := 1

	for k, v := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", k, argPos))
		args = append(args, v)
		argPos++
	}

	// Add updated_at timestamp
	setClauses = append(setClauses, fmt.Sprintf("updated_at = CURRENT_TIMESTAMP"))

	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id = $%d", strings.Join(setClauses, ", "), argPos)
	args = append(args, id)

	_, err := config.DB.Exec(query, args...)
	return err
}

func DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}

func FilterTasks( status *string, page int) ([]models.Task, int, error) {
	limit := 10
	offset := (page - 1) * limit

	query := `SELECT id, title, description, status, user_name, created_at, updated_at FROM tasks`
	var conditions []string
	var args []interface{}
	argPos := 1


	if status != nil {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argPos))
		args = append(args, *status)
		argPos++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY id ASC LIMIT %d OFFSET %d", limit, offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID, &task.Title, &task.Description, &task.Status,
			&task.UserName, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, task)
	}

	// Count total matching records
	countQuery := `SELECT COUNT(*) FROM tasks`
	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	row := config.DB.QueryRow(countQuery, args...)
	var total int
	if err := row.Scan(&total); err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}



func GetTaskByID(id string) (*models.Task, error) {
    query := `SELECT id, title, description, status, user_name, created_at, updated_at
              FROM tasks
              WHERE id = $1`
    var task models.Task
    row := config.DB.QueryRow(query, id)
    err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.UserName, &task.CreatedAt, &task.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &task, nil
}


