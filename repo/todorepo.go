package repo

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (tr *TodoRepository) Insert(todo *Todo) (*Todo, error) {
	tx, err := tr.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var t = Todo{}
	sql := `INSERT INTO todo (description, created_at, deleted) 
            VALUES ($1, $2, $3) 
            RETURNING *`

	err = tx.QueryRow(sql, &todo.Description, time.Now(), false).Scan(&t)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (tr *TodoRepository) GetByID(id int64) *Todo {
	return nil
}

func (tr *TodoRepository) GetByPage(page int, size int) []Todo {
	return nil
}

func (tr *TodoRepository) Update(todo *Todo) (*Todo, error) {
	tx, err := tr.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	sql := `UPDATE employee
            SET 
                description = $1, 
                updated_at = $2
            WHERE id = $3
            RETURNING *`

	t := Todo{}
	err = tx.QueryRow(
		sql, &todo.Description, time.Now(), &todo.ID).Scan(&t)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (tr *TodoRepository) Delete(id int64) (*Todo, error) {
	tx, err := tr.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	sql := `UPDATE employee
            SET deleted = $1
            WHERE id = $2
            RETURNING *`

	t := Todo{}
	err = tx.QueryRow(sql, false, id).Scan(&t)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &t, nil
}
