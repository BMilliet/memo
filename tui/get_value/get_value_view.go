package get_value

import tea "github.com/charmbracelet/bubbletea"

// View for getting a value
type GetValueView struct{}

func NewGetValueView() GetValueView {
	return GetValueView{}
}

func (g GetValueView) Init() tea.Cmd {
	return nil
}

func (g GetValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g GetValueView) View() string {
	return "Get Value View\n\nPress 'q' to quit"
}
