package logs

import (
    "unicode/utf8"
    "strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        switch char {
        case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
    }
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var modifiedLog strings.Builder
	for _, char := range log {
		if char == oldRune {
			modifiedLog.WriteString(string(newRune))
		} else {
			modifiedLog.WriteString(string(char))
		}
	}
	return modifiedLog.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}