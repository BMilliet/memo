package list_todos

import (
	"fmt"
	"memo/tui/interfaces"
	"memo/tui/styles"
	handler "memo/tui/todo_handler"
	utils "memo/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type TodosListView struct {
	mainView interfaces.MainViewInterface
	choices  []string       // items on the to-do list
	cursor   int            // which to-do list item our cursor is pointing at
	selected map[int]string // which to-do items are selected
}

func NewTodosListView(main interfaces.MainViewInterface) TodosListView {
	choices, err := handler.ReadExistingTodos()
	if err != nil {
		utils.LogMsg("Could not read existing todos")
	}

	return TodosListView{
		mainView: main,
		choices:  choices,
		selected: make(map[int]string),
	}
}

func (m TodosListView) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m TodosListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			updatedTodos := genUpdatedList(m)
			handler.SaveOverwriteTodos(updatedTodos)
			m.mainView.Quit()
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = m.choices[m.cursor]
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m TodosListView) View() string {
	// The header
	s := styles.TitleStyle.Render("\nWhich todos have you completed?:\n")

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
			checked = "x" // selected!
		}

		// This logic should be improved
		if cursor == ">" {
			s += styles.SelectedItemStyle.Render(fmt.Sprintf("\n%s [%s] %s", cursor, checked, choice))
		} else {
			s += fmt.Sprintf("\n%s [%s] %s", cursor, checked, choice)
		}
	}

	s += styles.FooterMsgStyle.Render("\n\nPress q to quit.\n")
	return s
}

func genUpdatedList(m TodosListView) []string {
	updatedList := removeSelected(m.choices, m.selected)
	return updatedList
}

func removeSelected(originalList []string, selectedList map[int]string) []string {
	filteredList := []string{}

	for i, item := range originalList {
		if _, selected := selectedList[i]; !selected {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}
