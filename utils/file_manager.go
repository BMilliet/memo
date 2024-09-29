package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	HomeDir string
	MemoDir string
}

func NewFileManager() (*FileManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %v", err)
	}

	memoDir := filepath.Join(homeDir, ".memo")
	return &FileManager{HomeDir: homeDir, MemoDir: memoDir}, nil
}

func (m *FileManager) ensureMemoDir() error {
	if _, err := os.Stat(m.MemoDir); os.IsNotExist(err) {
		err := os.Mkdir(m.MemoDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating .memo directory: %v", err)
		}
		fmt.Println("Created .memo directory")
	} else {
		fmt.Println(".memo directory already exists")
	}
	return nil
}

func (m *FileManager) checkAndCreateFile(filename string) error {
	filePath := filepath.Join(m.MemoDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating %s: %v", filename, err)
		}
		fmt.Printf("Created %s file\n", filename)
	} else {
		fmt.Printf("%s already exists\n", filename)
	}
	return nil
}

func (m *FileManager) BasicSetup() error {
	if err := m.ensureMemoDir(); err != nil {
		return err
	}

	files := []string{"todos.json", "saves.json"}
	for _, file := range files {
		if err := m.checkAndCreateFile(file); err != nil {
			return err
		}
	}

	return nil
}
