package src

import (
	"fmt"

	"github.com/google/uuid"
)

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
	snippets := "snippets"

	choices := []ListItem{
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
	case snippets:
		r.snippetsListSection()
	case temp:
		r.tempListSection()
	}
}

func (r *Runner) todoListSection() {
	// TODO: need to improve.
	todos := r.db.FindAllTodos()
	toRemove := r.viewBuilder.NewTodoListView(todos, 16)
	fmt.Println(toRemove)
}

func (r *Runner) snippetsListSection() {
	snippetsLists := r.db.FindAllSnippetsLists()
	if len(snippetsLists) == 0 {
		r.addSnippetList()
	}

	snippetsLists = r.db.FindAllSnippetsLists()
	choices := r.utils.CreateSnippetLists(snippetsLists)

	section := r.viewBuilder.NewListView("Select a snippet section.", choices, 16)

	switch section.OP {
	case "add":
		fmt.Println(section.OP)
		r.addSnippetList()
		return
	case "del":
		r.db.DeleteSnippetsList(section.ID)
		fmt.Println(section.OP)
		return
	}

	r.snippetsSection(section.ID)
}

func (r *Runner) snippetsSection(id string) {
	snippetItems := r.db.FindSnippetsByList(id)
	if len(snippetItems) == 0 {
		fmt.Println("add to empty section")
		r.addSnippetToList(id)
	}
	snippetItems = r.db.FindSnippetsByList(id)
	snippets := r.utils.ConvertSnippetItems(snippetItems)
	section := r.viewBuilder.NewListView("Select an snippet.", snippets, 16)

	switch section.OP {
	case "add":
		fmt.Println(section.OP)
		r.addSnippetToList(id)
		return
	case "del":
		r.db.DeleteSnippet(section.ID)
		fmt.Println(section.OP)
		return
	}

	fmt.Println(section.Content)
}

func (r *Runner) addSnippetList() {
	section := r.viewBuilder.NewTextFieldView("Write the section name for this snippet", "")

	snippetList := SnippetsList{
		ID:   uuid.New().String(),
		Name: section,
	}

	r.db.CreateSnippetsList(&snippetList)
}

func (r *Runner) addSnippetToList(id string) {
	name := r.viewBuilder.NewTextFieldView("Write the name for this snippet", "")
	content := r.viewBuilder.NewTextAreaFieldView("Write the snippet", "")

	snippet := Snippet{
		ID:             uuid.NewString(),
		SnippetsListID: id,
		Title:          name,
		Content:        content,
	}

	r.db.CreateSnippet(&snippet)
}

func (r *Runner) tempListSection() {
	fmt.Println("🚧 to implement")
}
