package src

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type textAreaInputViewModel struct {
	textInput textarea.Model
	err       error
	question  string
	endValue  *string
	quitting  bool
	styles    *Styles
	errors    bool
}

func TextAreaFieldViewModel(question, placeHolder string, value *string) textAreaInputViewModel {
	ti := textarea.New()
	ti.Placeholder = placeHolder
	ti.Focus()
	ti.CharLimit = 156
	ti.Placeholder = placeHolder

	return textAreaInputViewModel{
		textInput: ti,
		err:       nil,
		question:  question,
		endValue:  value,
		styles:    DefaultStyles(),
		quitting:  false,
	}
}

func (m textAreaInputViewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textAreaInputViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			v := m.textInput.Value()

			if v == "" {
				m.errors = true
				return m, cmd
			}

			*m.endValue = v
			m.quitting = true
			return m, tea.Quit

		case tea.KeyCtrlC, tea.KeyEsc:
			*m.endValue = ExitSignal
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textAreaInputViewModel) View() string {
	if m.quitting {
		return ""
	}

	inputField := m.styles.InputField.Render(m.textInput.View())
	if m.errors {
		inputField = m.styles.InputFieldWithError.Render(m.textInput.View())
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.TitleStyle.Render(fmt.Sprintf("\n%s\n", m.question)),
		inputField,
		m.styles.FooterStyle.Render("\n(ctrl+c or esc to quit)"),
	)
}

func TextAreaFieldView(title, placeHolder string, endValue *string) {
	m := TextAreaFieldViewModel(title, placeHolder, endValue)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("TextAreaFieldView -> ", err)
		os.Exit(1)
	}
}
