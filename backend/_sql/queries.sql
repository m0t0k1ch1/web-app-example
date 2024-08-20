-- name: CreateTask :execlastid
INSERT INTO task (title, updated_at, created_at) VALUES (?, ?, ?);

-- name: CountAllTasks :one
SELECT COUNT(*) FROM task;

-- name: CountTasksByStatus :one
SELECT COUNT(*) FROM task WHERE status = ?;

-- name: GetTask :one
SELECT * FROM task WHERE id = ?;

-- name: GetTaskForUpdate :one
SELECT * FROM task WHERE id = ? FOR UPDATE;

-- name: ListTasks :many
SELECT * FROM task ORDER BY id LIMIT ? OFFSET ?;

-- name: ListTasksByStatus :many
SELECT * FROM task WHERE status = ? ORDER BY id LIMIT ? OFFSET ?;

-- name: CompleteTask :exec
UPDATE task SET status = 'COMPLETED', updated_at = ? WHERE id = ?;
