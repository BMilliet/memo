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
	// Init and setup
	// Create instance of FileManager and setup.
	// FileManager should create the following:
	//
	// ~/.memo

	todos := r.db.FindAllTodos()
	fmt.Print(todos)
}
