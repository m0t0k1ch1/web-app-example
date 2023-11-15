-- name: CreateTask :execlastid
INSERT INTO task (title, updated_at, created_at) VALUES (?, ?, ?);

-- name: GetTask :one
SELECT * FROM task WHERE id = ?;

-- name: GetTaskForUpdate :one
SELECT * FROM task WHERE id = ? FOR UPDATE;

-- name: ListTasks :many
SELECT * FROM task ORDER BY id;

-- name: UpdateTaskTitle :exec
UPDATE task SET title = ?, updated_at = ? WHERE id = ?;

-- name: UpdateTaskStatus :exec
UPDATE task SET status = ?, updated_at = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM task WHERE id = ?;
