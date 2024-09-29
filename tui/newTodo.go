package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// View for adding a TODO
type addTodoView struct {
	question    Question
	width       int
	height      int
	answerField textinput.Model
	styles      *Styles
	todos       []string
}

type Question struct {
	question string
	answer   string
	// input    Input
}

func newAddTodoView() *addTodoView {
	styles := DefaultStyles()

	answerField := textinput.New()
	answerField.Placeholder = "Todo title"
	answerField.Focus()

	question := Question{question: "Todo title"}
	return &addTodoView{question: question, answerField: answerField, styles: styles}
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
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			return m, tea.Quit
		case "enter":
			m.question.answer = m.answerField.Value()
			m.next()
			m.answerField.SetValue("")
			return m, nil
		}
	}

	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

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

func (m addTodoView) View() string {
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
}
