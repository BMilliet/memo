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

func ParseJSONContent[T any](jsonString string) (*T, error) {
	var targetStruct T
	err := json.Unmarshal([]byte(jsonString), &targetStruct)
	if err != nil {
		return nil, fmt.Errorf("ParseJSONContent -> %v", err)
	}
	return &targetStruct, nil
}
