package list_todos

import tea "github.com/charmbracelet/bubbletea"

// View for listing TODOs
type ListTodoView struct{}

func NewListTodoView() ListTodoView {
	return ListTodoView{}
}

func (l ListTodoView) Init() tea.Cmd {
	return nil
}

func (l ListTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return l, tea.Quit
		}
	}
	return l, nil
}

func (l ListTodoView) View() string {
	return "List TODOs View\n\nPress 'q' to quit"
}
