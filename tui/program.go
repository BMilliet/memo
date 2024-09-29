package tui

import (
	// tea "github.com/charmbracelet/bubbletea"
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
		currentView: InitMenu(), // Start with the menu
	}
}

func (p *program) SwitchView(choice string) {
	switch choice {
	case "add TODO":
		p.currentView = newAddTodoView()
	case "list TODOs":
		p.currentView = newListTodoView()
	case "save value":
		p.currentView = newSaveValueView()
	case "get value":
		p.currentView = newGetValueView()
	}
}
