package src

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileManagerInterface interface {
	BasicSetup() error
}

type FileManager struct {
	homeDir        string
	memoDir        string
	DBPath         string
	MigrationsPath string
}

func NewFileManager() (*FileManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("NewFileManager -> %v", err)
	}

	memoDir := filepath.Join(homeDir, MemoDirName)
	memoDB := filepath.Join(memoDir, MemoDB)
	migrationsPath := filepath.Join(memoDir, MigrationsPath)

	return &FileManager{
		homeDir:        homeDir,
		memoDir:        memoDir,
		DBPath:         memoDB,
		MigrationsPath: migrationsPath,
	}, nil
}

func (m *FileManager) ensureMemoDir() error {
	if _, err := os.Stat(m.memoDir); os.IsNotExist(err) {
		err := os.Mkdir(m.memoDir, 0o755)
		if err != nil {
			return fmt.Errorf("ensureMemoDir -> %v", err)
		}
	}
	return nil
}

func (m *FileManager) BasicSetup() error {
	if err := m.ensureMemoDir(); err != nil {
		return err
	}

	return nil
}
