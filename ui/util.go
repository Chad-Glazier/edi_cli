package ui

import "strings"

// Repeat a string some number of times.
func Repeat(times int, str string) string {
	var s strings.Builder
	for range times {
		s.WriteString(str)
	}
	return s.String()
}
