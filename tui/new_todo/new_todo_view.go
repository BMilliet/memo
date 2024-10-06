package new_todo

import (
	"memo/tui/interfaces"
	styles "memo/tui/styles"
	handler "memo/tui/todo_handler"

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
	mainView    interfaces.MainViewInterface
	state       State
	question    Question
	width       int
	height      int
	answerField textinput.Model
	confirmList list.Model
	todos       []string
}

type Question struct {
	question string
	answer   string
}

func NewAddTodoView(main interfaces.MainViewInterface) *AddTodoView {

	answerField := textinput.New()
	answerField.Placeholder = "Todo title"
	answerField.Focus()

	question := Question{question: "Todo title"}

	// Initialize the confirmation list with "Yes" and "No"
	items := []list.Item{
		styles.Item("Yes"),
		styles.Item("No"),
	}
	confirmList := styles.ApplyStyle(items, "Add another TODO?")

	return &AddTodoView{
		mainView:    main,
		state:       stateAddTodo,
		question:    question,
		answerField: answerField,
		confirmList: confirmList,
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
			m.mainView.Quit()
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
				selected, ok := m.confirmList.SelectedItem().(styles.Item)
				if ok && selected == "Yes" {
					// User selected "Yes", return to add TODO
					m.state = stateAddTodo
					m.question.question = "Todo title"
				} else if ok && selected == "No" {
					handler.SaveNewTodos(m.todos)
					m.mainView.Quit()
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
}

// Render the view based on the current state
func (m AddTodoView) View() string {
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
				styles.InputField.Render(
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
