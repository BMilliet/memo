package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Initialize the program with the selected view
func (p *program) Init() tea.Cmd {
	return p.currentView.Init()
}

// Update logic that handles switching between views
func (p *program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Update the current view
	var cmd tea.Cmd
	p.currentView, cmd = p.currentView.Update(msg)

	// If the current view is the menu and a choice is made, switch views
	if m, ok := p.currentView.(model); ok && m.choice != "" {
		p.SwitchView(m.choice)
	}

	return p, cmd
}

// Render the current view
func (p *program) View() string {
	return p.currentView.View()
}

// Start the program
func StartTea() {
	p := NewProgram()
	if _, err := tea.NewProgram(p).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
