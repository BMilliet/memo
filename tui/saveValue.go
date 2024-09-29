package tui

import tea "github.com/charmbracelet/bubbletea"

// View for saving a value
type saveValueView struct{}

func newSaveValueView() ViewModel {
	return saveValueView{}
}

func (s saveValueView) Init() tea.Cmd {
	return nil
}

func (s saveValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return s, tea.Quit
		}
	}
	return s, nil
}

func (s saveValueView) View() string {
	return "Save Value View\n\nPress 'q' to quit"
}
