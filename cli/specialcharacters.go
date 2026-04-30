package cli

import "strings"

const (
	LINE_HORIZONTAL     = "\u2500" // ─
	LINE_VERTICAL       = "\u2502" // │
	CORNER_TOP_LEFT     = "\u250C" // ┌
	CORNER_TOP_RIGHT    = "\u2510" // ┐
	CORNER_BOTTOM_LEFT  = "\u2514" // └
	CORNER_BOTTOM_RIGHT = "\u2518" // ┘
	T_VERTICAL_RIGHT    = "\u251C" // ├
	T_VERTICAL_LEFT     = "\u2524" // ┤
	T_HORIZONTAL_DOWN   = "\u252C" // ┬
	T_HORIZONTAL_UP     = "\u2534" // ┴
	CROSS               = "\u253C" // ┼
	DIAG_UPPER_RIGHT    = "\u2571" // ╱
	DIAG_LOWER_LEFT     = "\u2571" // ╱
	DIAG_UPPER_LEFT     = "\u2572" // ╲
	DIAG_LOWER_RIGHT    = "\u2572" // ╲
)

// Repeat a string some number of times.
func repeat(times int, str string) string {
	var s strings.Builder
	for range times {
		s.WriteString(str)
	}
	return s.String()
}
