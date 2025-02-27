package main

import (
	"fmt"
	"log"
	"os"

	"memo/src"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Println(src.MemoVersion)
			return
		}
	}

	fileManager, err := src.NewFileManager()
	if err != nil {
		log.Fatalln(err, "Failed to initialize FileManager")
	}

	utils := src.NewUtils()
	viewBuilder := src.NewViewBuilder()

	runner := src.NewRunner(
		fileManager,
		utils,
		viewBuilder,
	)

	runner.Start()
}
