package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	logMsgStyle   = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	errorMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

func LogMsg(format string, args ...interface{}) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	styledLogMsg := logMsgStyle.Render(format)
	fmt.Println(styledLogMsg)
}

func ErrorMsg(format string, args ...interface{}) string {
	errorMsg := fmt.Sprintf(format, args...)
	styledErrorMsg := errorMsgStyle.Render(errorMsg)
	return styledErrorMsg
}
