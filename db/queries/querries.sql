-- name: CreateUser :one
INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
    username = coalesce(sqlc.narg(username), username),
    password = coalesce(sqlc.narg(password), password),
    email = coalesce(sqlc.narg(email), email)
WHERE username = sqlc.arg(username)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: CreateGroup :one
INSERT INTO groups (name, description) VALUES ($1, $2) RETURNING *;

-- name: GetGroup :one
SELECT * FROM groups WHERE id = $1 LIMIT 1;

-- name: ListGroups :many
SELECT * FROM groups LIMIT $1 OFFSET $2;

-- name: UpdateGroup :one
UPDATE groups
SET
    name = coalesce(sqlc.narg(name), name),
    description = coalesce(sqlc.narg(description), description)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM groups WHERE id = $1;

-- name: AddUserToGroup :exec
INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2);

-- name: RemoveUserFromGroup :exec
DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2;

-- name: CreateTask :one
INSERT INTO tasks (title, description, due_date, status, group_id, user_id) VALUES ($1, $2, $3, $4, $5, $6)RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks;

-- name: UpdateTask :one
UPDATE tasks
SET
    title = coalesce(sqlc.narg(title), title),
    description = coalesce(sqlc.narg(description), description),
    due_date = coalesce(sqlc.narg(due_date), due_date),
    status = coalesce(sqlc.narg(status), status)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: CreateSubtask :one
INSERT INTO subtasks (title, description, due_date, status, task_id) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetSubtask :one
SELECT * FROM subtasks WHERE id = $1 LIMIT 1;

-- name: ListTaskSubtasks :many
SELECT * FROM subtasks WHERE task_id = $1;

-- name: UpdateSubtask :one
UPDATE subtasks
SET
    title = coalesce(sqlc.narg(title), title),
    description = coalesce(sqlc.narg(description), description),
    due_date = coalesce(sqlc.narg(due_date), due_date),
    status = coalesce(sqlc.narg(status), status)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteSubtask :exec
DELETE FROM subtasks WHERE id = $1;