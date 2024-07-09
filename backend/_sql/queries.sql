-- name: CreateTask :execlastid
INSERT INTO task (title, updated_at, created_at) VALUES (?, ?, ?);

-- name: GetTask :one
SELECT * FROM task WHERE id = ?;

-- name: GetTaskForUpdate :one
SELECT * FROM task WHERE id = ? FOR UPDATE;

-- name: ListFirstTasks :many
SELECT * FROM task ORDER BY id LIMIT ?;

-- name: ListFirstTasksAfterCursor :many
SELECT * FROM task WHERE id > ? ORDER BY id LIMIT ?;

-- name: ListFirstTasksByStatus :many
SELECT * FROM task WHERE status = ? ORDER BY id LIMIT ?;

-- name: ListFirstTasksAfterCursorByStatus :many
SELECT * FROM task WHERE status = ? AND id > ? ORDER BY id LIMIT ?;

-- name: CompleteTask :exec
UPDATE task SET status = 'COMPLETED', updated_at = ? WHERE id = ?;
