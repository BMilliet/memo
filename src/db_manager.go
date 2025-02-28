package src

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DBManagerInterface interface {
	Setup()

	// To-Do CRUD
	CreateTodo(todo *Todo) Todo
	GetTodo(id string) *Todo
	FindAllTodos() []*Todo
	UpdateTodo(todo *Todo)
	DeleteTodo(id string)

	// Snippets List CRUD
	CreateSnippetsList(snippetsList *SnippetsList) SnippetsList
	GetSnippetsList(id string) *SnippetsList
	FindAllSnippetsLists() []*SnippetsList
	DeleteSnippetsList(id string)

	// Snippet CRUD
	CreateSnippet(snippet *Snippet) Snippet
	GetSnippet(id string) *Snippet
	FindSnippetsByList(snippetsListID string) []*Snippet
	UpdateSnippet(snippet *Snippet)
	DeleteSnippet(id string)
}

type DBManager struct {
	db      *sql.DB
	ctx     context.Context
	queries *Queries
	utils   UtilsInterface
	dbPath  string
}

func NewDbManager(utils UtilsInterface, dbPath string) *DBManager {
	return &DBManager{
		utils:  utils,
		dbPath: dbPath,
	}
}

func (dbm *DBManager) Setup() {
	fmt.Println("🔌 Connecting to db...")

	db, err := sql.Open("sqlite3", dbm.dbPath)
	if err != nil {
		dbm.utils.HandleError(err, "failed to open database")
	}

	dbm.db = db
	dbm.queries = New(db)
	fmt.Println("✅ Database setup complete.")
}

// To-Do CRUD

func (dbm *DBManager) CreateTodo(todo *Todo) Todo {
	ctx := context.Background()
	obj, err := dbm.queries.CreateTodo(ctx, CreateTodoParams{
		ID:    todo.ID,
		Title: todo.Title,
	})
	dbm.utils.HandleError(err, "CreateTodo")
	return obj
}

func (dbm *DBManager) GetTodo(id string) *Todo {
	ctx := context.Background()
	todo, err := dbm.queries.GetTodo(ctx, id)
	dbm.utils.HandleError(err, "GetTodo")

	return &Todo{
		ID:    todo.ID,
		Title: todo.Title,
	}
}

func (dbm *DBManager) FindAllTodos() []*Todo {
	ctx := context.Background()
	results, err := dbm.queries.SelectAllTodos(ctx)
	dbm.utils.HandleError(err, "FindAllTodos")

	var todos []*Todo
	for _, result := range results {
		todos = append(todos, &Todo{
			ID:    result.ID,
			Title: result.Title,
		})
	}
	return todos
}

func (dbm *DBManager) UpdateTodo(todo *Todo) {
	ctx := context.Background()
	err := dbm.queries.UpdateTodo(ctx, UpdateTodoParams{
		Title: todo.Title,
		ID:    todo.ID,
	})
	dbm.utils.HandleError(err, "UpdateTodo")
}

func (dbm *DBManager) DeleteTodo(id string) {
	ctx := context.Background()
	err := dbm.queries.DeleteTodo(ctx, id)
	dbm.utils.HandleError(err, "DeleteTodo")
}

// Snippets List CRUD

func (dbm *DBManager) CreateSnippetsList(snippetsList *SnippetsList) SnippetsList {
	ctx := context.Background()
	obj, err := dbm.queries.CreateSnippetsList(ctx, CreateSnippetsListParams{
		ID:   snippetsList.ID,
		Name: snippetsList.Name,
	})
	dbm.utils.HandleError(err, "CreateSnippetsList")
	return obj
}

func (dbm *DBManager) GetSnippetsList(id string) *SnippetsList {
	ctx := context.Background()
	list, err := dbm.queries.GetSnippetsList(ctx, id)
	dbm.utils.HandleError(err, "GetSnippetsList")

	return &SnippetsList{
		ID:   list.ID,
		Name: list.Name,
	}
}

func (dbm *DBManager) FindAllSnippetsLists() []*SnippetsList {
	ctx := context.Background()
	results, err := dbm.queries.SelectAllSnippetsLists(ctx)
	dbm.utils.HandleError(err, "FindAllSnippetsLists")

	var lists []*SnippetsList
	for _, result := range results {
		lists = append(lists, &SnippetsList{
			ID:   result.ID,
			Name: result.Name,
		})
	}
	return lists
}

func (dbm *DBManager) DeleteSnippetsList(id string) {
	ctx := context.Background()
	err := dbm.queries.DeleteSnippetsList(ctx, id)
	dbm.utils.HandleError(err, "DeleteSnippetsList")
}

// Snippet CRUD

func (dbm *DBManager) CreateSnippet(snippet *Snippet) Snippet {
	ctx := context.Background()
	obj, err := dbm.queries.CreateSnippet(ctx, CreateSnippetParams{
		ID:             snippet.ID,
		SnippetsListID: snippet.SnippetsListID,
		Title:          snippet.Title,
		Content:        snippet.Content,
	})
	dbm.utils.HandleError(err, "CreateSnippet")
	return obj
}

func (dbm *DBManager) GetSnippet(id string) *Snippet {
	ctx := context.Background()
	snippet, err := dbm.queries.GetSnippet(ctx, id)
	dbm.utils.HandleError(err, "GetSnippet")

	return &Snippet{
		ID:             snippet.ID,
		SnippetsListID: snippet.SnippetsListID,
		Title:          snippet.Title,
		Content:        snippet.Content,
	}
}

func (dbm *DBManager) FindSnippetsByList(snippetsListID string) []*Snippet {
	ctx := context.Background()
	results, err := dbm.queries.ListSnippetsByList(ctx, snippetsListID)
	dbm.utils.HandleError(err, "FindSnippetsByList")

	var snippets []*Snippet
	for _, result := range results {
		snippets = append(snippets, &Snippet{
			ID:             result.ID,
			SnippetsListID: result.SnippetsListID,
			Title:          result.Title,
			Content:        result.Content,
		})
	}
	return snippets
}

func (dbm *DBManager) UpdateSnippet(snippet *Snippet) {
	ctx := context.Background()
	err := dbm.queries.UpdateSnippet(ctx, UpdateSnippetParams{
		Title:   snippet.Title,
		Content: snippet.Content,
		ID:      snippet.ID,
	})
	dbm.utils.HandleError(err, "UpdateSnippet")
}

func (dbm *DBManager) DeleteSnippet(id string) {
	ctx := context.Background()
	err := dbm.queries.DeleteSnippet(ctx, id)
	dbm.utils.HandleError(err, "DeleteSnippet")
}
