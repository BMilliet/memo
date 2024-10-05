package values_handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type LabelValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// ReadLabelValues reads and returns the list of label-value pairs from the JSON file.
func ReadLabelValues() ([]LabelValue, error) {
	filePath, err := getSavedValuesPath()
	if err != nil {
		return nil, fmt.Errorf("failed to read saved file path")
	}

	var labelValues []LabelValue

	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		// Read the existing content
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read existing JSON file: %v", err)
		}

		// Unmarshal the existing JSON content into labelValues slice
		if len(fileData) > 0 {
			if err := json.Unmarshal(fileData, &labelValues); err != nil {
				return nil, fmt.Errorf("failed to unmarshal existing JSON content: %v", err)
			}
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("error checking if file exists: %v", err)
	}

	return labelValues, nil
}

// merging the new pair with existing data if the file already contains values.
func WriteLabelValue(label string, value string) error {
	filePath, err := getSavedValuesPath()
	if err != nil {
		return fmt.Errorf("failed to read saved file path")
	}
	// Read existing label-value pairs from the file
	labelValues, err := ReadLabelValues()
	if err != nil {
		return err
	}

	// Add the new label-value pair to the list
	newLabelValue := LabelValue{
		Label: label,
		Value: value,
	}
	labelValues = append(labelValues, newLabelValue)

	// Marshal the updated list back to JSON
	updatedData, err := json.MarshalIndent(labelValues, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal updated JSON: %v", err)
	}

	// Write the updated JSON back to the file
	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write to JSON file: %v", err)
	}

	return nil
}

func getSavedValuesPath() (string, error) {
	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %v", err)
	}

	memoDir := filepath.Join(homeDir, ".memo")
	todoFilePath := filepath.Join(memoDir, "saved.json")
	return todoFilePath, nil
}
