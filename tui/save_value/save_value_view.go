package save_value

import (
	"memo/tui/interfaces"
	"memo/tui/styles"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	defaultWidth = 20
	listHeight   = 14
)

type ViewModel struct {
}

// View for saving a value
type SaveValueView struct {
	mainView      interfaces.MainViewInterface
	questions     []Question
	questionIndex int
}

type Question struct {
	question  string
	answer    string
	textField Input
}

func NewSaveValueView(main interfaces.MainViewInterface) *SaveValueView {
	questions := []Question{
		{question: "Insert label", textField: NewShortAnswerField()},
		{question: "Copy your content", textField: NewLongAnswerField()},
	}

	answerField := textinput.New()
	answerField.Focus()

	return &SaveValueView{
		mainView:      main,
		questions:     questions,
		questionIndex: 0,
	}
}

func (m SaveValueView) Init() tea.Cmd {
	return nil
}

func (m *SaveValueView) next() {
	if m.questionIndex < len(m.questions)-1 {
		m.questionIndex++
	} else {
		// someting else
		m.mainView.Quit()
	}
}

func (m SaveValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	currentQuestion := &m.questions[m.questionIndex]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.mainView.Quit()
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			currentQuestion.answer = currentQuestion.textField.Value()
			m.next()
			return m, nil
		}
	}

	currentQuestion.textField, cmd = currentQuestion.textField.Update(msg)
	return m, cmd
}

func (m SaveValueView) View() string {
	current := m.questions[m.questionIndex]

	return lipgloss.Place(
		defaultWidth,
		listHeight,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			current.question,
			styles.InputField.Render(
				current.textField.View(),
			),
		),
	)
}
