package src

import (
	"encoding/json"
	"fmt"
	"os"
)

type UtilsInterface interface {
	ValidateInput(v string)
	HandleError(err error, message string)
	ExitWithError(message string)
	CreateSnippetLists(snippets []*SnippetsList) []ListItem
	ConvertSnippetItems(snippets []*Snippet) []ListItem
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

func ParseJSONContent[T any](jsonString string) (*T, error) {
	var targetStruct T
	err := json.Unmarshal([]byte(jsonString), &targetStruct)
	if err != nil {
		return nil, fmt.Errorf("ParseJSONContent -> %v", err)
	}
	return &targetStruct, nil
}
