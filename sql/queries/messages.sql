-- name: CreateMessage :one
INSERT INTO messages (id, chat_id, role, text, created_at) VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: GetMessagesByChatID :many
SELECT * FROM messages WHERE chat_id = ? ORDER BY created_at;
