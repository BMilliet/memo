package src

import (
	"fmt"
	"os"
	"os/exec"
)

type UtilsInterface interface {
	ValidateInput(v string)
	HandleError(err error, message string)
	ExitWithError(message string)
	CreateSnippetLists(snippets []*SnippetsList) []ListItem
	ConvertSnippetItems(snippets []*Snippet) []ListItem
	CopyToClipboard(text string) error
}

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) ValidateInput(v string) {
	if v == ExitSignal {
		os.Exit(0)
	}
}

func (u *Utils) HandleError(err error, message string) {
	if err != nil {
		msg := fmt.Sprintf(message+" -> ", err)
		u.ExitWithError(msg)
	}
}

func (u *Utils) ExitWithError(message string) {
	s := DefaultStyles()
	println((s.Text(message, s.ErrorColor)))
	os.Exit(1)
}

func (u *Utils) CreateSnippetLists(snippets []*SnippetsList) []ListItem {
	var items []ListItem

	for _, snippet := range snippets {
		item := ListItem{
			ID: snippet.ID,
			T:  snippet.Name,
			D:  snippet.ID,
		}
		items = append(items, item)
	}

	return items
}

func (u *Utils) ConvertSnippetItems(snippets []*Snippet) []ListItem {
	var items []ListItem

	for _, snippet := range snippets {
		item := ListItem{
			ID:      snippet.ID,
			T:       snippet.Title,
			D:       u.truncate(snippet.Content, 25),
			Content: snippet.Content,
		}
		items = append(items, item)
	}

	return items
}

func (u *Utils) truncate(s string, max int) string {
	if len(s) > max {
		return s[:max]
	}
	return s
}

func (u *Utils) CopyToClipboard(text string) error {
	cmd := exec.Command("pbcopy")

	in, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start pbcopy: %v", err)
	}

	_, err = in.Write([]byte(text))
	if err != nil {
		return fmt.Errorf("failed to write to pbcopy: %v", err)
	}

	if err := in.Close(); err != nil {
		return fmt.Errorf("failed to close stdin pipe: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait for pbcopy: %v", err)
	}

	return nil
}
