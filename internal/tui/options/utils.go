package options

import (
	"fmt"
	"strconv"
	"strings"
)

func GetEmojiByCleanOption(optionName string) string {
	emoji := ""

	switch optionName {
	case ShowHiddenFiles:
		emoji = "👁️"
	case ConfirmDeletion:
		emoji = "⚠️"
	case IncludeSubfolders:
		emoji = "📁"
	case DeleteEmptySubfolders:
		emoji = "🗑️"
	case SendFilesToTrash:
		emoji = "♻️"
	case SecureDeleteFiles:
		emoji = "🔒"
	case LogOperations:
		emoji = "📝"
	case LogToFile:
		emoji = "📄"
	case ShowStatistics:
		emoji = "📊"
	case ExitAfterDeletion:
		emoji = "🚪"
	}

	return emoji
}

// GetNextOption returns the next or previous option in a circular manner
func GetNextOption(currentOption, optionPrefix string, maxOptions int, forward bool) string {
	currentNum := 1
	if strings.HasPrefix(currentOption, optionPrefix) {
		numStr := strings.TrimPrefix(currentOption, optionPrefix)
		if num, err := strconv.Atoi(numStr); err == nil {
			currentNum = num
		}
	}

	var nextNum int
	if forward {
		nextNum = currentNum + 1
		if nextNum > maxOptions {
			nextNum = 1
		}
	} else {
		nextNum = currentNum - 1
		if nextNum < 1 {
			nextNum = maxOptions
		}
	}

	return fmt.Sprintf(optionPrefix + fmt.Sprintf("%d", nextNum))
}
