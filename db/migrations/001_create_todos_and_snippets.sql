-- +goose Up
CREATE TABLE todos (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE snippets_list (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE snippets (
    id TEXT PRIMARY KEY,
    snippets_list_id TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (snippets_list_id) REFERENCES snippets_list(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE snippets;
DROP TABLE snippets_list;
DROP TABLE todos;

