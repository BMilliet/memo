package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// var (
// 	//Program    *tea.Program
// 	Program    *program
// 	WindowSize tea.WindowSizeMsg
// )

// Define the struct for your main program
// type program struct {
// 	currentView ViewModel
// }

// Styles
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render
var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render
