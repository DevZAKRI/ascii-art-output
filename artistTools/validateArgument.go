package output

import (
	"os"
	"strings"
)

// ValidateArgument validates the command-line arguments for the ASCII art tool.
// It checks if the number of arguments is correct and if the input string contains valid ASCII characters.
// If the arguments are not valid, it returns an error message.
// If the arguments are valid, it returns an empty string.
func ValidateArgument(args []string) (string, bool) {
	if len(args) < 2 || len(args) > 4 {
		return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard", false
	}
	isFlag := false
	if strings.HasPrefix(args[1], "--output=") {
		isFlag = true
	}
	if isFlag {
		if len(args[1]) < 14 || !strings.HasSuffix(args[1], ".txt") {
			return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard", isFlag
		}
		if !IsValidASCII(args[2]) {
			return "STRING Containe Invalid CHARACTERS", isFlag
		}
	} else {
		if !IsValidASCII(args[1]) {
			return "STRING Containe Invalid CHARACTERS", isFlag
		}
	}
	if !isFlag && len(args) == 4 {
		return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard", isFlag
	}
	if len(os.Args) == 4 && !IsValidBanner(os.Args[3]) {
		return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard", isFlag
	}
	return "", isFlag
}

func IsValidBanner(args string) bool {
	if args == "standard" || args == "shadow" || args == "thinkertoy" || args == "standard.txt" || args == "shadow.txt" || args == "thinkertoy.txt" {
		return true
	}
	return false
}
