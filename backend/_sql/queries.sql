-- name: CreateTask :execlastid
INSERT INTO task (title, updated_at, created_at) VALUES (?, ?, ?);

-- name: GetTask :one
SELECT * FROM task WHERE id = ?;

-- name: GetTaskForUpdate :one
SELECT * FROM task WHERE id = ? FOR UPDATE;

-- name: ListTasks :many
SELECT * FROM task ORDER BY id;

-- name: ListTasksByStatus :many
SELECT * FROM task WHERE status = ? ORDER BY id;

-- name: UpdateTask :exec
UPDATE task SET title = ?, status = ?, updated_at = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM task WHERE id = ?;
