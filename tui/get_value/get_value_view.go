package get_value

import (
	"memo/tui/interfaces"

	tea "github.com/charmbracelet/bubbletea"
)

// View for getting a value
type GetValueView struct {
	mainView interfaces.MainViewInterface
}

func NewGetValueView(main interfaces.MainViewInterface) GetValueView {
	return GetValueView{mainView: main}
}

func (g GetValueView) Init() tea.Cmd {
	return nil
}

func (m GetValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			m.mainView.Quit()
			return m, tea.Quit
		}
	}
	return m, nil
}

func (g GetValueView) View() string {
	return "Get Value View\n\nPress 'q' to quit"
}
