package utils

import (
	"fmt"
	"os/exec"
)

// CopyToClipboard copies the provided text to the macOS clipboard using pbcopy
func CopyToClipboard(text string) error {
	// Create an exec.Command to run pbcopy
	cmd := exec.Command("pbcopy")

	// Get the input pipe of the pbcopy command
	in, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %v", err)
	}

	// Start the pbcopy command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start pbcopy: %v", err)
	}

	// Write the text to the pbcopy input
	_, err = in.Write([]byte(text))
	if err != nil {
		return fmt.Errorf("failed to write to pbcopy: %v", err)
	}

	// Close the input pipe
	if err := in.Close(); err != nil {
		return fmt.Errorf("failed to close stdin pipe: %v", err)
	}

	// Wait for pbcopy to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait for pbcopy: %v", err)
	}

	return nil
}
