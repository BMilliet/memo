package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// View for adding a TODO
type addTodoView struct{}

func newAddTodoView() ViewModel {
	return addTodoView{}
}

func (a addTodoView) Init() tea.Cmd {
	return nil
}

func (a addTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return a, tea.Quit
		}
	}
	return a, nil
}

func (a addTodoView) View() string {
	return "Add TODO View\n\nPress 'q' to quit"
}
