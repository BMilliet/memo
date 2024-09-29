package tui

import tea "github.com/charmbracelet/bubbletea"

// View for getting a value
type getValueView struct{}

func newGetValueView() ViewModel {
	return getValueView{}
}

func (g getValueView) Init() tea.Cmd {
	return nil
}

func (g getValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g getValueView) View() string {
	return "Get Value View\n\nPress 'q' to quit"
}
