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
	HomeDir string
	MemoDir string
	MemoDB  string
}

func NewFileManager() (*FileManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("NewFileManager -> %v", err)
	}

	memoDir := filepath.Join(homeDir, MemoDirName)
	memoDB := filepath.Join(memoDir, MemoDB)

	return &FileManager{
		HomeDir: homeDir,
		MemoDir: memoDir,
		MemoDB:  memoDB,
	}, nil
}

func (m *FileManager) ensureMemoDir() error {
	if _, err := os.Stat(m.MemoDir); os.IsNotExist(err) {
		err := os.Mkdir(m.MemoDir, 0o755)
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
