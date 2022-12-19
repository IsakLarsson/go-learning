package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	characters := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}
	for _, char := range log {
		value, exists := characters[char]
		if exists {
			return value
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, char := range log {
		if char == oldRune {
			char = newRune
		}
		newLog += string(char)
	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
