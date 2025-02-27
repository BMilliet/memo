package src

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileManagerInterface interface {
	CheckIfPathExists(path string) (bool, error)
	ReadFileContent(filePath string) (string, error)
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
		err := os.Mkdir(m.MemoDB, 0o755)
		if err != nil {
			return fmt.Errorf("ensureMemoDir -> %v", err)
		}
	}
	return nil
}

func (m *FileManager) CheckIfPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("CheckIfPathExists -> %v", err)
}

func (m *FileManager) checkAndCreateFile(filePath string) error {
	exists, err := m.CheckIfPathExists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("checkAndCreateFile -> %s %v", filePath, err)
		}
	}
	return nil
}

func (m *FileManager) ReadFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("ReadFileContent -> %s %v", filePath, err)
	}
	return string(data), nil
}

func (m *FileManager) GetDBContent() (string, error) {
	str, err := m.ReadFileContent(m.MemoDB)
	if err != nil {
		return "", fmt.Errorf("GetDBContent -> %s %v", m.MemoDB, err)
	}
	return str, nil
}

func (m *FileManager) WriteDBContent(content string) error {
	err := m.writeFileContent(m.MemoDB, content)
	if err != nil {
		return fmt.Errorf("WriteDBContent -> %s: %v", m.MemoDB, err)
	}
	return nil
}

func (m *FileManager) writeFileContent(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("writeFileContent -> %s %v", filePath, err)
	}
	return nil
}

func (m *FileManager) BasicSetup() error {
	if err := m.ensureMemoDir(); err != nil {
		return err
	}

	files := []string{
		m.MemoDB,
	}

	for _, file := range files {
		if err := m.checkAndCreateFile(file); err != nil {
			return err
		}
	}

	return nil
}
