-- name: GetById :one
SELECT * FROM employee
WHERE id = $1 LIMIT 1;

-- name: GetAll :many
SELECT * FROM employee;

-- name: Insert :exec
INSERT INTO employee (name, age, created_at)
VALUES ($1, $2, $3);

-- name: Update :execrows
Update employee
SET
    name = $1,
    age = $2
WHERE
    id = $3;