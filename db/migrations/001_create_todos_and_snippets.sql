
-- +goose Up
CREATE TABLE todos (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE snippets_list (
    id TEXT PRIMARY KEY DEFAULT,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE snippets (
    id TEXT PRIMARY KEY DEFAULT,
    session_id TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (snippets_list_id) REFERENCES sessions(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE snippets;
DROP TABLE snippets_list;
DROP TABLE todo_list;

