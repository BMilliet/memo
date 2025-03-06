// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package src

import (
	"context"
)

const createSnippet = `-- name: CreateSnippet :one
INSERT INTO snippets (id, snippets_list_id, title, content) 
VALUES (?, ?, ?, ?)
RETURNING id, snippets_list_id, title, content, created_at
`

type CreateSnippetParams struct {
	ID             string
	SnippetsListID string
	Title          string
	Content        string
}

func (q *Queries) CreateSnippet(ctx context.Context, arg CreateSnippetParams) (Snippet, error) {
	row := q.db.QueryRowContext(ctx, createSnippet,
		arg.ID,
		arg.SnippetsListID,
		arg.Title,
		arg.Content,
	)
	var i Snippet
	err := row.Scan(
		&i.ID,
		&i.SnippetsListID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const createSnippetsList = `-- name: CreateSnippetsList :one
INSERT INTO snippets_list (id, name) 
VALUES (?, ?)
RETURNING id, name, created_at
`

type CreateSnippetsListParams struct {
	ID   string
	Name string
}

func (q *Queries) CreateSnippetsList(ctx context.Context, arg CreateSnippetsListParams) (SnippetsList, error) {
	row := q.db.QueryRowContext(ctx, createSnippetsList, arg.ID, arg.Name)
	var i SnippetsList
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (id, title) 
VALUES (?, ?)
RETURNING id, title, created_at
`

type CreateTodoParams struct {
	ID    string
	Title string
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.ID, arg.Title)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

const deleteSnippet = `-- name: DeleteSnippet :exec
DELETE FROM snippets WHERE id = ?
`

func (q *Queries) DeleteSnippet(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteSnippet, id)
	return err
}

const deleteSnippetsList = `-- name: DeleteSnippetsList :exec
DELETE FROM snippets_list WHERE id = ?
`

func (q *Queries) DeleteSnippetsList(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteSnippetsList, id)
	return err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getSnippet = `-- name: GetSnippet :one
SELECT id, snippets_list_id, title, content, created_at FROM snippets WHERE id = ?
`

func (q *Queries) GetSnippet(ctx context.Context, id string) (Snippet, error) {
	row := q.db.QueryRowContext(ctx, getSnippet, id)
	var i Snippet
	err := row.Scan(
		&i.ID,
		&i.SnippetsListID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getSnippetsList = `-- name: GetSnippetsList :one
SELECT id, name, created_at FROM snippets_list WHERE id = ?
`

func (q *Queries) GetSnippetsList(ctx context.Context, id string) (SnippetsList, error) {
	row := q.db.QueryRowContext(ctx, getSnippetsList, id)
	var i SnippetsList
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, created_at FROM todos WHERE id = ?
`

func (q *Queries) GetTodo(ctx context.Context, id string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.CreatedAt)
	return i, err
}

const listSnippetsByList = `-- name: ListSnippetsByList :many
SELECT id, snippets_list_id, title, content, created_at FROM snippets WHERE snippets_list_id = ? ORDER BY created_at DESC
`

func (q *Queries) ListSnippetsByList(ctx context.Context, snippetsListID string) ([]Snippet, error) {
	rows, err := q.db.QueryContext(ctx, listSnippetsByList, snippetsListID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Snippet
	for rows.Next() {
		var i Snippet
		if err := rows.Scan(
			&i.ID,
			&i.SnippetsListID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSnippetsLists = `-- name: ListSnippetsLists :many
SELECT id, name, created_at FROM snippets_list ORDER BY created_at DESC
`

func (q *Queries) ListSnippetsLists(ctx context.Context) ([]SnippetsList, error) {
	rows, err := q.db.QueryContext(ctx, listSnippetsLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SnippetsList
	for rows.Next() {
		var i SnippetsList
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, created_at FROM todos ORDER BY created_at DESC
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectAllSnippets = `-- name: SelectAllSnippets :many
SELECT id, snippets_list_id, title, content, created_at FROM snippets
`

func (q *Queries) SelectAllSnippets(ctx context.Context) ([]Snippet, error) {
	rows, err := q.db.QueryContext(ctx, selectAllSnippets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Snippet
	for rows.Next() {
		var i Snippet
		if err := rows.Scan(
			&i.ID,
			&i.SnippetsListID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectAllSnippetsLists = `-- name: SelectAllSnippetsLists :many
SELECT id, name, created_at FROM snippets_list
`

func (q *Queries) SelectAllSnippetsLists(ctx context.Context) ([]SnippetsList, error) {
	rows, err := q.db.QueryContext(ctx, selectAllSnippetsLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SnippetsList
	for rows.Next() {
		var i SnippetsList
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectAllTodos = `-- name: SelectAllTodos :many
SELECT id, title, created_at FROM todos
`

func (q *Queries) SelectAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, selectAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSnippet = `-- name: UpdateSnippet :exec
UPDATE snippets SET title = ?, content = ? WHERE id = ?
`

type UpdateSnippetParams struct {
	Title   string
	Content string
	ID      string
}

func (q *Queries) UpdateSnippet(ctx context.Context, arg UpdateSnippetParams) error {
	_, err := q.db.ExecContext(ctx, updateSnippet, arg.Title, arg.Content, arg.ID)
	return err
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos SET title = ? WHERE id = ?
`

type UpdateTodoParams struct {
	Title string
	ID    string
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.ExecContext(ctx, updateTodo, arg.Title, arg.ID)
	return err
}
