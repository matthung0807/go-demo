package repo

import (
	"database/sql"

	"abc.com/demo/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return EmployeeRepository{
		db: db,
	}
}

func (er *EmployeeRepository) GetAllEmployees() ([]model.Employee, error) {
	rows, err := er.db.Query("SELECT id, name, age, created_at FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []model.Employee
	for rows.Next() {
		var e model.Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Age, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		emps = append(emps, e)
	}
	return emps, nil
}

func (er *EmployeeRepository) GetEmployeeByID(db *sql.DB, id int64) (*model.Employee, error) {
	row := er.db.QueryRow("SELECT * FROM employee WHERE id = $1 LIMIT 1", id)
	var emp model.Employee
	err := row.Scan(
		&emp.ID,
		&emp.Name,
		&emp.Age,
		&emp.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}
