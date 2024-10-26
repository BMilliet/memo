package menu

import (
	"memo/tui/interfaces"
	"memo/tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the interface for all views (models)
type ViewModel interface {
	tea.Model
}

// Main controller model to handle view switching
type MainView struct {
	currentView ViewModel
	quitting    bool
	customMsg   []string
}

var _ interfaces.MainViewInterface = (*MainView)(nil)

func (p *MainView) Quit(messages ...string) {
	p.customMsg = messages
	p.quitting = true
}

func NewProgram() *MainView {
	return &MainView{
		currentView: initMenu(), // Start with the menu
		quitting:    false,
	}
}

// Initialize the program with the selected view
func (p *MainView) Init() tea.Cmd {
	return p.currentView.Init()
}

// Update logic that handles switching between views
func (p *MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Update the current view
	var cmd tea.Cmd
	p.currentView, cmd = p.currentView.Update(msg)

	if p.quitting {
		return p, tea.Quit
	}

	// If the current view is the menu and a choice is made, switch views
	if m, ok := p.currentView.(MenuViewModel); ok && m.choice != "" {
		p.SwitchView(m.choice)
	}

	return p, cmd
}

// Render the current view
func (p *MainView) View() string {
	if p.quitting {
		message := "See ya 👋 💾"

		if len(p.customMsg) > 0 {
			message += " " + p.customMsg[0]
		}

		return styles.QuitTextStyle.Render(message)
	}
	return p.currentView.View()
}
