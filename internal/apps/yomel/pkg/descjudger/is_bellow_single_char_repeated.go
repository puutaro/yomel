package descjudger

import (
	"strings"
)

func IsBellowSingleCharRepeated(s string) bool {
	trimmed := strings.Trim(s, " 　")
	if len(trimmed) == 0 {
		return true
	}
	runes := []rune(strings.ToLower(trimmed))
	seen := make(map[rune]bool)
	for _, r := range runes {
		seen[r] = true
		if len(seen) > 1 {
			return false
		}
	}
	return len(seen) == 1
}
