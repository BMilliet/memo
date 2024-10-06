package menu

import (
	"memo/tui/styles"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuViewModel struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m MenuViewModel) Init() tea.Cmd {
	return nil
}

func (m MenuViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(styles.Item)
			if ok {
				m.choice = string(i) // Set the choice to trigger view change
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MenuViewModel) View() string {
	if m.quitting {
		return styles.QuitTextStyle.Render("See ya 👋")
	}
	return "\n" + m.list.View()
}

// Initialize the menu with items
func initMenu() MenuViewModel {
	items := []list.Item{
		styles.Item("list TODOs"),
		styles.Item("add TODO"),
		styles.Item("save value"),
		styles.Item("get value"),
	}

	l := styles.ApplyStyle(items, "Welcome to memo 💾")
	return MenuViewModel{list: l}
}
