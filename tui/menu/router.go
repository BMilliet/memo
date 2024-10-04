package menu

import (
	getValue "memo/tui/get_value"
	listTodos "memo/tui/list_todos"
	addTodo "memo/tui/new_todo"
	saveValue "memo/tui/save_value"
)

func (p *MainView) SwitchView(choice string) {
	switch choice {
	case "add TODO":
		p.currentView = addTodo.NewAddTodoView(p)
	case "list TODOs":
		p.currentView = listTodos.NewTodosListView(p)
	case "save value":
		p.currentView = saveValue.NewSaveValueView(p)
	case "get value":
		p.currentView = getValue.NewGetValueView(p)
	}
}
