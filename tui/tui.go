package tui

import (
	"fmt"
	main "memo/tui/menu"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Start the program
func StartTea() {
	p := main.NewProgram()
	if _, err := tea.NewProgram(p).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
