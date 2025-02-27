-- name: CreateTodo :one
INSERT INTO todos (id, title, completed) 
VALUES (?, ?, ?)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = ?;

-- name: ListTodos :many
SELECT * FROM todos ORDER BY created_at DESC;

-- name: SelectAllTodos :many
SELECT * FROM todos;

-- name: UpdateTodo :exec
UPDATE todos SET title = ?, completed = ? WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;

-- name: CreateSnippetsList :one
INSERT INTO snippets_list (id, name) 
VALUES (?, ?)
RETURNING *;

-- name: GetSnippetsList :one
SELECT * FROM snippets_list WHERE id = ?;

-- name: ListSnippetsLists :many
SELECT * FROM snippets_list ORDER BY created_at DESC;

-- name: SelectAllSnippetsLists :many
SELECT * FROM snippets_list;

-- name: DeleteSnippetsList :exec
DELETE FROM snippets_list WHERE id = ?;

-- name: CreateSnippet :one
INSERT INTO snippets (id, snippets_list_id, title, content) 
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetSnippet :one
SELECT * FROM snippets WHERE id = ?;

-- name: ListSnippetsByList :many
SELECT * FROM snippets WHERE snippets_list_id = ? ORDER BY created_at DESC;

-- name: SelectAllSnippets :many
SELECT * FROM snippets;

-- name: UpdateSnippet :exec
UPDATE snippets SET title = ?, content = ? WHERE id = ?;

-- name: DeleteSnippet :exec
DELETE FROM snippets WHERE id = ?;
