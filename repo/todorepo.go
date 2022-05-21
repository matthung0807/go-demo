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
type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (tr *TodoRepositoryImpl) Insert(todo *Todo) (*Todo, error) {
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

func (tr *TodoRepositoryImpl) GetByID(id int64) (*Todo, error) {
	sql := "SELECT * FROM todo WHERE id = $1"
	row := tr.db.QueryRow(sql, id)
	var todo Todo
	err := row.Scan(
		&todo.ID,
		&todo.Description,
		&todo.CreatedAt,
		&todo.Deleted,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (tr *TodoRepositoryImpl) GetByPage(page int, size int) ([]Todo, int, error) {
	tx, err := tr.db.Begin()
	if err != nil {
		return nil, 0, err
	}
	sql := "SELECT FROM todo LIMIT $1 OFFSET $2"
	rows, err := tx.Query(sql, size, page-1)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Description,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		todos = append(todos, todo)
	}

	sql = "SELECT count(*) FROM todo"
	c := 0
	err = tx.QueryRow(sql).Scan(&c)
	if err != nil {
		return nil, 0, err
	}

	return todos, c, nil
}

func (tr *TodoRepositoryImpl) Update(todo *Todo) (*Todo, error) {
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

func (tr *TodoRepositoryImpl) Delete(id int64) (*Todo, error) {
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
