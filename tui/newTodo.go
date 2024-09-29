package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// View for adding a TODO
type addTodoView struct{}

func newAddTodoView() ViewModel {
	return addTodoView{}
}

func (a addTodoView) Init() tea.Cmd {
	return nil
}

func (a addTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return a, tea.Quit
		}
	}
	return a, nil
}

func (a addTodoView) View() string {
	return "Add TODO View\n\nPress 'q' to quit"
}

// View for listing TODOs
type listTodoView struct{}

func newListTodoView() ViewModel {
	return listTodoView{}
}

func (l listTodoView) Init() tea.Cmd {
	return nil
}

func (l listTodoView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return l, tea.Quit
		}
	}
	return l, nil
}

func (l listTodoView) View() string {
	return "List TODOs View\n\nPress 'q' to quit"
}

// View for saving a value
type saveValueView struct{}

func newSaveValueView() ViewModel {
	return saveValueView{}
}

func (s saveValueView) Init() tea.Cmd {
	return nil
}

func (s saveValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return s, tea.Quit
		}
	}
	return s, nil
}

func (s saveValueView) View() string {
	return "Save Value View\n\nPress 'q' to quit"
}

// View for getting a value
type getValueView struct{}

func newGetValueView() ViewModel {
	return getValueView{}
}

func (g getValueView) Init() tea.Cmd {
	return nil
}

func (g getValueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g getValueView) View() string {
	return "Get Value View\n\nPress 'q' to quit"
}
