-- name: CreateTask :execlastid
INSERT INTO `tasks` (`title`) VALUES (?);

-- name: GetTask :one
SELECT * FROM `tasks` WHERE `id` = ?;

-- name: GetTaskForUpdate :one
SELECT * FROM `tasks` WHERE `id` = ? FOR UPDATE;

-- name: ListTasks :many
SELECT * FROM `tasks` ORDER BY `id` DESC;

-- name: UpdateTask :exec
UPDATE `tasks` SET `title` = ?, `is_completed` = ?, `updated_at` = UNIX_TIMESTAMP(NOW()) WHERE `id` = ?;

-- name: DeleteTask :exec
DELETE FROM `tasks` WHERE `id` = ?;
