-- messages table
CREATE TABLE messages (
    id TEXT PRIMARY KEY,
    chat_id TEXT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK (role IN ('user', 'assistant', 'system', 'developer', 'tool')),
    text TEXT NOT NULL,
    created_at INTEGER NOT NULL
);

-- index for ordering messages
CREATE INDEX idx_messages_chat_id_created_at ON messages(chat_id, created_at);
