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
SELECT * FROM task
WHERE
  CASE WHEN CAST(sqlc.arg(set_status) AS UNSIGNED) > 0
    THEN status = sqlc.arg(status)
    ELSE 1
  END
ORDER BY id LIMIT ? OFFSET ?;

-- name: CompleteTask :exec
UPDATE task SET status = 'COMPLETED', updated_at = ? WHERE id = ?;
