package src

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type TextReturnObject struct {
	OP, Content string
}

type textInputViewModel struct {
	textInput textinput.Model
	err       error
	question  string
	endValue  *TextReturnObject
	quitting  bool
	styles    *Styles
	errors    bool
}

func TextFieldViewModel(question, placeHolder string, textReturnObj *TextReturnObject) textInputViewModel {
	ti := textinput.New()
	ti.Placeholder = placeHolder
	ti.Focus()
	ti.CharLimit = 255
	ti.Placeholder = placeHolder

	return textInputViewModel{
		textInput: ti,
		err:       nil,
		question:  question,
		endValue:  textReturnObj,
		styles:    DefaultStyles(),
		quitting:  false,
	}
}

func (m textInputViewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textInputViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

			*m.endValue = TextReturnObject{OP: AddSignal, Content: v}
			m.quitting = true
			return m, tea.Quit

		case tea.KeyCtrlC, tea.KeyEsc:
			*m.endValue = TextReturnObject{OP: ExitSignal, Content: ""}
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textInputViewModel) View() string {
	if m.quitting {
		message := "See ya 👋 💾"
		return message
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

func TextFieldView(title, placeHolder string, endValue *TextReturnObject) {
	m := TextFieldViewModel(title, placeHolder, endValue)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("TextFieldView -> ", err)
		os.Exit(1)
	}
}
