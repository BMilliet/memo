package new_todo

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// States to manage the different views
type State int

const (
	stateAddTodo State = iota // Adding TODO view
	stateConfirm              // Confirmation view
)

// View for adding a TODO
type AddTodoView struct {
	state       State
	question    Question
	width       int
	height      int
	answerField textinput.Model
	confirmList list.Model
	styles      *Styles
	todos       []string
	quitting    bool
}

type Question struct {
	question string
	answer   string
}

func NewAddTodoView() *AddTodoView {
	styles := DefaultStyles()

	answerField := textinput.New()
	answerField.Placeholder = "Todo title"
	answerField.Focus()

	const defaultWidth = 20
	const listHeight = 14

	question := Question{question: "Todo title"}

	// Initialize the confirmation list with "Yes" and "No"
	items := []list.Item{
		item("Yes"),
		item("No"),
	}
	confirmList := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	confirmList.Title = "Add another TODO?"

	confirmList.SetShowStatusBar(false)
	confirmList.SetFilteringEnabled(false)
	confirmList.Styles.Title = titleStyle
	confirmList.Styles.PaginationStyle = paginationStyle
	confirmList.Styles.HelpStyle = helpStyle

	return &AddTodoView{
		state:       stateAddTodo,
		question:    question,
		answerField: answerField,
		confirmList: confirmList,
		styles:      styles,
	}
}

func (a AddTodoView) Init() tea.Cmd {
	return nil
}

func (m AddTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.state == stateAddTodo {
				// Save the TODO and switch to the confirmation view
				m.question.answer = m.answerField.Value()
				m.next()
				m.answerField.SetValue("")
				m.state = stateConfirm // Switch to confirmation view
			} else if m.state == stateConfirm {
				// Check the user's selection
				selected, ok := m.confirmList.SelectedItem().(item)
				if ok && selected == "Yes" {
					// User selected "Yes", return to add TODO
					m.state = stateAddTodo
					m.question.question = "Todo title"
				} else if ok && selected == "No" {
					m.quitting = true
					return m, tea.Quit
				}
			}
		}
	}

	// Update the list or input field based on the current state
	if m.state == stateAddTodo {
		m.answerField, cmd = m.answerField.Update(msg)
	} else if m.state == stateConfirm {
		m.confirmList, cmd = m.confirmList.Update(msg)
	}

	return m, cmd
}

// Add the new TODO and log it to file
func (m *AddTodoView) next() {
	m.todos = append(m.todos, m.question.answer)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	} else {
		// IMPLEMENT HERE
		fmt.Fprintf(f, "todo: %s\n", m.todos[len(m.todos)-1])
	}
	defer f.Close()
}

// Render the view based on the current state
func (m AddTodoView) View() string {
	if m.quitting {
		return quitTextStyle.Render("See ya 👋")
	}

	switch m.state {
	case stateAddTodo:
		// Render the add TODO view
		current := m.question
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				current.question,
				m.styles.InputField.Render(
					m.answerField.View(),
				),
			),
		)

	case stateConfirm:
		// Render the confirmation view (Yes/No list)
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			m.confirmList.View(),
		)
	}
	return ""
}
