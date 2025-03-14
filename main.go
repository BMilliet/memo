package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"memo/src"

	"github.com/pressly/goose/v3"
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

	// Init and setup
	// Create instance of FileManager and setup.
	// FileManager should create the following:
	//
	// ~/.memo
	setupErr := fileManager.BasicSetup()
	if setupErr != nil {
		log.Fatalln(setupErr, "Failed to setup basic memo dotfiles")
	}

	dbPath := fileManager.DBPath
	migrationsPath := fileManager.MigrationsPath

	setupDB(dbPath, migrationsPath)

	utils := src.NewUtils()
	viewBuilder := src.NewViewBuilder()
	db := src.NewDbManager(utils, dbPath)

	db.Setup()

	runner := src.NewRunner(
		fileManager,
		utils,
		viewBuilder,
		db,
	)

	runner.Start()
}

func setupDB(path string, migrationPath string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	if err := runMigrations(db, migrationPath); err != nil {
		log.Fatal("Migration failed:", err)
	}
}

func runMigrations(db *sql.DB, migrationsDir string) error {
	goose.SetDialect("sqlite3")

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}
	return nil
}
