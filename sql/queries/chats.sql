-- name: CreateChat :one
INSERT INTO chats (id, name) VALUES (?, ?) RETURNING *;

-- name: GetChats :many
SELECT * FROM chats ORDER BY id;

-- name: GetChat :one
SELECT * FROM chats WHERE id = ?;