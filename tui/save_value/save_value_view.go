package save_value

import tea "github.com/charmbracelet/bubbletea"

// View for saving a value
type SaveValueView struct{}

func NewSaveValueView() SaveValueView {
	return SaveValueView{}
}

func (s SaveValueView) Init() tea.Cmd {
	return nil
}

func (s SaveValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return s, tea.Quit
		}
	}
	return s, nil
}

func (s SaveValueView) View() string {
	return "Save Value View\n\nPress 'q' to quit"
}
