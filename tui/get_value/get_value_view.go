package get_value

import (
	"fmt"
	"memo/tui/interfaces"
	handle "memo/tui/values_handler"
	"memo/utils"

	tea "github.com/charmbracelet/bubbletea"
)

// View for getting a value
type GetValueView struct {
	mainView interfaces.MainViewInterface
	choices  []handle.LabelValue // items on the to-do list
	cursor   int                 // which to-do list item our cursor is pointing at
	selected map[int]struct{}
}

func NewGetValueView(main interfaces.MainViewInterface) GetValueView {
	saved, _ := handle.ReadLabelValues()

	return GetValueView{
		mainView: main,
		choices:  saved,
		selected: make(map[int]struct{}),
	}
}

func (g GetValueView) Init() tea.Cmd {
	return nil
}

func (m GetValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			m.mainView.Quit()
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				selectedValue := m.choices[m.cursor].Value
				utils.CopyToClipboard(selectedValue)
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m GetValueView) View() string {
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "✅"
		}

		// Render the row
		s += fmt.Sprintf("%s %s %s\n", cursor, checked, choice.Label)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
