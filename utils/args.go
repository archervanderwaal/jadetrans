package utils

import (
	"strings"
)

// ParseArgs are defined for parsing parameters,
// and returns values are the sentences or words and query parameters to be queried.
func ParseArgs(osArgs []string) ([]string, []string) {
	words := make([]string, 0)
	args := []string{osArgs[0]}
	lastArg := ""
	for _, arg := range osArgs[1:] {
		if strings.HasPrefix(arg, "-") {
			args = append(args, arg)
			lastArg = arg
			continue
		}
		if strings.HasPrefix(lastArg, "-") && len(words) != 0 {
			continue
		}
		words = append(words, arg)
		lastArg = arg
	}
	return words, args[1:]
}
