package new_todo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(
		s.BorderColor,
	).BorderStyle(
		lipgloss.NormalBorder(),
	).Padding(1).Width(80)
	return s
}

// States to manage the different views
type state int

const (
	stateAddTodo state = iota // Adding TODO view
	stateConfirm              // Confirmation view
)

// View for adding a TODO
type addTodoView struct {
	state       state
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

func NewAddTodoView() *addTodoView {
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

	// Apply styles similar to menuView
	// confirmList.Styles.Title = titleStyle
	// confirmList.Styles.PaginationStyle = paginationStyle
	// confirmList.Styles.HelpStyle = helpStyle
	confirmList.SetShowStatusBar(false)
	confirmList.SetFilteringEnabled(false)
	confirmList.Styles.Title = titleStyle
	confirmList.Styles.PaginationStyle = paginationStyle
	confirmList.Styles.HelpStyle = helpStyle

	return &addTodoView{
		state:       stateAddTodo,
		question:    question,
		answerField: answerField,
		confirmList: confirmList,
		styles:      styles,
	}
}

func (a addTodoView) Init() tea.Cmd {
	return nil
}

func (m addTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
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
func (m *addTodoView) next() {
	m.todos = append(m.todos, m.question.answer)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(f, "todo: %s\n", m.todos[len(m.todos)-1])
	}
	defer f.Close()
}

// Render the view based on the current state
func (m addTodoView) View() string {
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

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
