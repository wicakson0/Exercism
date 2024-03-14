package logs

import "unicode/utf8"

// import "strings"

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
	// Note: strings are IMMUTABLE in Go, to modify convert to a slice of rune
	// or use the "strings" std library
	// solution 1: using "strings"
	// return strings.ReplaceAll(log, string(oldRune), string(newRune))

	// solution 2: convert to a slice of rune
	runeSlice := []rune(log)

	for index, char := range runeSlice {
		if char == oldRune {
			runeSlice[index] = newRune
		}
	}

	return string(runeSlice)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
