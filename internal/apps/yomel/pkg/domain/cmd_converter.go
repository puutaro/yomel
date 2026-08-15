package domain

import (
	"fmt"
	"strings"
	"unicode"
)

func makeEscapeComment(
	comment string,
	isTerminal bool,
) string {
	if comment == "" {
		return ""
	}
	commentoutWithSpaceBackQuote := fmt.Sprintf(
		"`%s%s`",
		SpaceCommentout,
		toLowerWithSpaces(comment),
	)
	if !isTerminal {
		return commentoutWithSpaceBackQuote
	}
	colorEnd := "\x1b[39m"
	return fmt.Sprintf(
		"%s%s%s",
		GrayStart,
		commentoutWithSpaceBackQuote,
		colorEnd,
	)
}

func toLowerWithSpaces(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	runesLen := len(runes)
	for i := 0; i < runesLen; i++ {
		r := runes[i]
		nextRuneIndex := i + 1
		if !unicode.IsUpper(r) ||
			(nextRuneIndex < runesLen &&
				unicode.IsUpper(runes[nextRuneIndex])) {
			sb.WriteRune(r)
			continue
		}
		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteRune(unicode.ToLower(r))
	}
	return sb.String()
}
