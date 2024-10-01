package utils

import (
	"fmt"
	"memo/tui/styles"
)

func LogMsg(format string, args ...interface{}) {
	logMessage("log", format, args...)
}

func ErrorMsg(format string, args ...interface{}) {
	logMessage("error", format, args...)
}

func logMessage(msgType string, format string, args ...interface{}) {
	var styledMsg string
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}

	switch msgType {
	case "log":
		styledMsg = styles.LogMsgStyle.Render(format)
	case "error":
		styledMsg = styles.ErrorMsgStyle.Render(format)
	default:
		styledMsg = format
	}

	fmt.Println(styledMsg)
}
