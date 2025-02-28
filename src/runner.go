package src

import "fmt"

type Runner struct {
	fileManager FileManagerInterface
	utils       UtilsInterface
	viewBuilder ViewBuilderInterface
	db          DBManagerInterface
}

func NewRunner(fm FileManagerInterface, u UtilsInterface, b ViewBuilderInterface, db DBManagerInterface) *Runner {
	return &Runner{
		fileManager: fm,
		utils:       u,
		viewBuilder: b,
		db:          db,
	}
}

func (r *Runner) Start() {
	temp := "temp"
	todo := "todos"
	snippets := "snippets"

	choices := []ListItem{
		{
			T: todo,
			D: "📝 Todo list",
		},
		{
			T: snippets,
			D: "⚡️ snippets list",
		},
		{
			T: temp,
			D: "⏳ temporary saved values",
		},
	}

	answer := r.viewBuilder.NewListView("Select one menu option.", choices, 16)
	r.utils.ValidateInput(answer.T)

	switch answer.T {
	case todo:
		r.todoListSection()
	case snippets:
		r.snippetsListSection()
	case temp:
		r.tempListSection()
	}
}

func (r *Runner) todoListSection() {
	todos := r.db.FindAllTodos()
	fmt.Println(todos)
}

func (r *Runner) snippetsListSection() {
	fmt.Println("🚧 to implement")
}

func (r *Runner) tempListSection() {
	fmt.Println("🚧 to implement")
}
