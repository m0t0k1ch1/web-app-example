-- name: CreateTask :execlastid
INSERT INTO task (title) VALUES (?);

-- name: GetTask :one
SELECT * FROM task WHERE id = ?;

-- name: GetTaskForUpdate :one
SELECT * FROM task WHERE id = ? FOR UPDATE;

-- name: ListTasks :many
SELECT * FROM task ORDER BY id DESC;

-- name: UpdateTask :exec
UPDATE task SET title = ?, status = ?, updated_at = UNIX_TIMESTAMP(NOW()) WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM task WHERE id = ?;
