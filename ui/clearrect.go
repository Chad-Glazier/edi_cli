package ui

import (
	"fmt"
	"strings"

	"github.com/Chad-Glazier/edi_cli/ansi"
)

// ClearRect clears the rectangular region of the terminal defined by the space
// between (inclusive) the two row number and the two column numbers
func ClearRect(row1, row2, col1, col2 int) {

	ansi.HideCursor()
	defer ansi.ShowCursor()

	startRow := min(row1, row2)
	startCol := min(col1, col2)
	endRow := max(row1, row2)
	endCol := max(col1, col2)

	blankLine := strings.Repeat(" ", endCol-startCol)

	for r := startRow; r <= endRow; r++ {
		ansi.SetCursor(r, startCol)
		fmt.Print(blankLine)
	}

}
