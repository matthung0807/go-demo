package repo

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
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
	sql := `INSERT INTO todo (description) 
            VALUES ($1, $2, $3) 
            RETURNING *`

	err = tx.QueryRow(sql, &todo.Description).
		Scan(
			&t.ID,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.Deleted,
		)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (tr *TodoRepositoryImpl) GetByID(id int64) (*Todo, error) {
	sql := `SELECT * FROM todo 
			WHERE id = $1 
				AND deleted = false`
	row := tr.db.QueryRow(sql, id)
	var todo Todo
	err := row.Scan(
		&todo.ID,
		&todo.Description,
		&todo.CreatedAt,
		&todo.UpdatedAt,
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
	sql := `SELECT * FROM todo 
			WHERE deleted = false 
			LIMIT $1 OFFSET $2`
	rows, err := tx.Query(sql, size, (page-1)*size)
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
			&todo.Deleted,
		)
		if err != nil {
			return nil, 0, err
		}
		todos = append(todos, todo)
	}

	sql = `SELECT count(*) 
			FROM todo 
			WHERE deleted = false`
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

	sql := `UPDATE todo
            SET 
                description = $1, 
                updated_at = $2
            WHERE id = $3
            RETURNING *`

	t := Todo{}
	err = tx.QueryRow(
		sql, &todo.Description, time.Now(), &todo.ID).
		Scan(
			&t.ID,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.Deleted,
		)

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

	sql := `UPDATE todo
            SET deleted = $1
            WHERE id = $2
            RETURNING *`

	t := Todo{}
	err = tx.QueryRow(sql, true, id).
		Scan(
			&t.ID,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.Deleted,
		)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &t, nil
}
