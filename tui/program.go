package tui

import (
	getValue "memo/tui/get_value"
	listTodo "memo/tui/list_todos"
	addTodo "memo/tui/new_todo"
	saveValue "memo/tui/save_value"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the interface for all views (models)
type ViewModel interface {
	tea.Model
}

// Main controller model to handle view switching
type program struct {
	currentView ViewModel
}

func NewProgram() *program {
	return &program{
		currentView: initMenu(), // Start with the menu
	}
}

func (p *program) SwitchView(choice string) {
	switch choice {
	case "add TODO":
		p.currentView = addTodo.NewAddTodoView()
	case "list TODOs":
		p.currentView = listTodo.NewListTodoView()
	case "save value":
		p.currentView = saveValue.NewSaveValueView()
	case "get value":
		p.currentView = getValue.NewGetValueView()
	}
}
