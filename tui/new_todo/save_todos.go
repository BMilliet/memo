package new_todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SaveTodos saves the new todos into ".memo/todo.json" after merging with existing ones
func saveTodos(newTodos []string) error {
	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	// Construct the file path
	memoDir := filepath.Join(homeDir, ".memo")
	todoFilePath := filepath.Join(memoDir, "todos.json")

	// Read the existing todos from "todo.json" file
	existingTodos, err := readExistingTodos(todoFilePath)
	if err != nil {
		return fmt.Errorf("failed to read existing todos: %w", err)
	}

	// Merge existing todos with the new ones
	mergedTodos := mergeTodos(existingTodos, newTodos)

	// Save the merged todos back to the "todo.json" file
	err = writeTodosToFile(todoFilePath, mergedTodos)
	if err != nil {
		return fmt.Errorf("failed to write todos to file: %w", err)
	}

	return nil
}

// readExistingTodos reads the existing todos from the "todo.json" file
func readExistingTodos(filePath string) ([]string, error) {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If file doesn't exist, return an empty slice (no todos exist yet)
		return []string{}, nil
	}

	// Read the content of the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the JSON content into a slice of strings
	var todos []string
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return todos, nil
}

// mergeTodos merges the existing todos with new todos, avoiding duplicates
func mergeTodos(existingTodos, newTodos []string) []string {
	todoMap := make(map[string]bool)

	// Add existing todos to the map
	for _, todo := range existingTodos {
		todoMap[todo] = true
	}

	// Add new todos to the map
	for _, todo := range newTodos {
		todoMap[todo] = true
	}

	// Convert the map back to a slice
	var mergedTodos []string
	for todo := range todoMap {
		mergedTodos = append(mergedTodos, todo)
	}

	return mergedTodos
}

// writeTodosToFile writes the merged todos into the "todo.json" file
func writeTodosToFile(filePath string, todos []string) error {
	// Marshal the todos into JSON format
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write the JSON data back to the file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
