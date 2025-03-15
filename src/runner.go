package src

import (
	"fmt"
	"os"

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

	answer := r.viewBuilder.NewListView("Select one menu option.", choices, 14, false)
	r.utils.ValidateInput(answer.T)

	switch answer.T {
	case snippets:
		r.snippetsListSection()
	case temp:
		r.tempListSection()
	}

	r.exit()
}

func (r *Runner) snippetsListSection() {
	snippetsLists := r.db.FindAllSnippetsLists()
	if len(snippetsLists) == 0 {
		r.addSnippetList()
	}

	snippetsLists = r.db.FindAllSnippetsLists()
	choices := r.utils.CreateSnippetLists(snippetsLists)

	section := r.viewBuilder.NewListView("Select a snippet section.", choices, 24, true)

	switch section.OP {
	case AddSignal:
		r.addSnippetList()
		return
	case RemoveSignal:
		r.db.DeleteSnippetsList(section.ID)
		return
	case ExitSignal:
		r.exit()
	}

	r.snippetsSection(section.ID)
}

func (r *Runner) snippetsSection(id string) {
	snippetItems := r.db.FindSnippetsByList(id)
	if len(snippetItems) == 0 {
		r.addSnippetToList(id)
	}
	snippetItems = r.db.FindSnippetsByList(id)
	snippets := r.utils.ConvertSnippetItems(snippetItems)
	section := r.viewBuilder.NewListView("Select an snippet.", snippets, 24, true)

	switch section.OP {
	case AddSignal:
		r.addSnippetToList(id)
		return
	case RemoveSignal:
		r.db.DeleteSnippet(section.ID)
		return
	case ExitSignal:
		r.exit()
	}

	r.copyToClipBoard(section.Content)
	r.exit()
}

func (r *Runner) addSnippetList() {
	section := r.viewBuilder.NewTextFieldView("Write the section name for this snippet", "")
	if section.OP == ExitSignal {
		r.exit()
	}

	snippetList := SnippetsList{
		ID:   uuid.New().String(),
		Name: section.Content,
	}

	r.db.CreateSnippetsList(&snippetList)
}

func (r *Runner) addSnippetToList(id string) {
	name := r.viewBuilder.NewTextFieldView("Write the name for this snippet", "")
	if name.OP == ExitSignal {
		r.exit()
	}

	content := r.viewBuilder.NewTextAreaFieldView("Write the snippet", "")
	if content.OP == ExitSignal {
		r.exit()
	}

	snippet := Snippet{
		ID:             uuid.NewString(),
		SnippetsListID: id,
		Title:          name.Content,
		Content:        content.Content,
	}

	r.db.CreateSnippet(&snippet)
}

func (r *Runner) exit() {
	fmt.Println("💾 See ya 👋")
	os.Exit(0)
}

func (r *Runner) copyToClipBoard(text string) {
	r.utils.CopyToClipboard(text)
	fmt.Println("\033[32m✅ Content copied to clipboard\033[0m")
}

func (r *Runner) tempListSection() {
	fmt.Println("🚧 to implement")
}
