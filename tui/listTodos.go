package tui

import tea "github.com/charmbracelet/bubbletea"

// View for listing TODOs
type listTodoView struct{}

func newListTodoView() ViewModel {
	return listTodoView{}
}

func (l listTodoView) Init() tea.Cmd {
	return nil
}

func (l listTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return l, tea.Quit
		}
	}
	return l, nil
}

func (l listTodoView) View() string {
	return "List TODOs View\n\nPress 'q' to quit"
}
