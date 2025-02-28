package src

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TodoListItem struct {
	T, OT, ID string
	S         bool
}

func (i TodoListItem) Title() string       { return i.T }
func (i TodoListItem) FilterValue() string { return i.T }

type TodoListViewModel struct {
	list     list.Model
	selected map[string]bool
	endValue *[]string
	quitting bool
	styles   Styles
}

func (m TodoListViewModel) Init() tea.Cmd {
	return nil
}

func (m TodoListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "enter":
			i, _ := m.list.SelectedItem().(TodoListItem)

			i.S = !i.S

			if i.S {
				i.T = i.T + " OK"
				m.selected[i.ID] = true
			} else {
				i.T = i.OT
				delete(m.selected, i.ID)
			}

			m.list.SetItems(m.list.Items())

		case "ctrl+c", "esc", "q":
			tmp := []string{}
			for k := range m.selected {
				tmp = append(tmp, k)
			}
			*m.endValue = tmp
			return m, tea.Quit

		case "a", "m", "d":
			// TODO: CRUD
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m TodoListViewModel) View() string {
	if m.quitting {
		return ""
	}

	return m.list.View()
}

func TodoListView(op []*Todo, height int, endValue *[]string) {
	title := "📝 Todo list"
	items := []list.Item{}
	for _, o := range op {
		item := TodoListItem{
			T:  o.Title,
			OT: o.Title,
			S:  false,
		}
		items = append(items, item)
	}

	styles := DefaultStyles()

	const defaultWidth = 20

	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = styles.SelectedItemStyle
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle.
		Foreground(styles.SelectedTitleColor)

	l := list.New(items, delegate, defaultWidth, height)
	l.Title = fmt.Sprintf("\n%s", title)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = styles.TitleStyle
	l.Styles.Title.Align(lipgloss.Left)
	l.Styles.PaginationStyle = styles.PaginationStyle
	l.Styles.HelpStyle = styles.HelpStyle

	m := TodoListViewModel{
		list:     l,
		endValue: endValue,
		selected: make(map[string]bool),
		styles:   *styles,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("TodoListView -> ", err)
		os.Exit(1)
	}
}
