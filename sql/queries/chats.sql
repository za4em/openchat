-- name: CreateChat :one
INSERT INTO chats (id, name, updated_at) VALUES (?, ?, ?) RETURNING *;

-- name: GetChats :many
SELECT * FROM chats ORDER BY id;

-- name: GetChat :one
SELECT * FROM chats WHERE id = ?;

-- name: UpdateChat :exec
UPDATE chats SET name = ?, updated_at = ? WHERE id = ?;

-- name: UpdateChatUpdatedAt :exec
UPDATE chats SET updated_at = ? WHERE id = ?;
