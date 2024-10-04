package save_value

import (
	"memo/tui/interfaces"

	tea "github.com/charmbracelet/bubbletea"
)

// View for saving a value
type SaveValueView struct {
	mainView interfaces.MainViewInterface
}

func NewSaveValueView(main interfaces.MainViewInterface) SaveValueView {
	return SaveValueView{mainView: main}
}

func (s SaveValueView) Init() tea.Cmd {
	return nil
}

func (m SaveValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			m.mainView.Quit()
			return m, tea.Quit
		}
	}
	return m, nil
}

func (s SaveValueView) View() string {
	return "Save Value View\n\nPress 'q' to quit"
}
