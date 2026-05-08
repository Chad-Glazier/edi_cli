package out

import (
	"fmt"

	"github.com/Chad-Glazier/edi/bb"
	"github.com/Chad-Glazier/edi/state"
	"github.com/Chad-Glazier/edi_cli/ansi"
)

// Prints a given board state to the standard output. The startRow and startCol
// define the location of the top-left corner of the board; the returned values
// specify the bottom-right corner of the board.
func PrintState(
	board state.Board, startRow, startCol int,
) (
	endRow, endCol int,
) {
	ansi.HideCursor()
	defer ansi.ShowCursor()

	r := startRow
	c := startCol

	ansi.SetCursor(r, c)
	fmt.Print("    0 1 2 3 4 5 6 7 8 9 ")

	r++
	ansi.SetCursor(r, c)
	fmt.Print("  " +
		ansi.CORNER_TOP_LEFT +
		ansi.Repeat(21, ansi.LINE_HORIZONTAL) +
		ansi.CORNER_TOP_RIGHT,
	)

	for row := range 10 {
		r++
		ansi.SetCursor(r, c)

		fmt.Printf("%d %s", row, ansi.LINE_VERTICAL)

		for col := 0; col < 10; col++ {
			s := ansi.FgBrightBlack(ansi.VACANT_SQUARE)

			switch {
			case board.White.Flagged(bb.Pos(row, col)):
				s = ansi.FgBrightCyan(ansi.WHITE_QUEEN)
				if board.Player == state.WHITE {
					s = ansi.Blink(s)
				}
			case board.Black.Flagged(bb.Pos(row, col)):
				s = ansi.FgBrightRed(ansi.BLACK_QUEEN)
				if board.Player == state.BLACK {
					s = ansi.Blink(s)
				}
			case board.Occupancy.Flagged(bb.Pos(row, col)):
				s = ansi.FgWhite(ansi.ARROW_SQUARE)
			}

			fmt.Print(" " + s)
		}

		fmt.Print(" " + ansi.LINE_VERTICAL)
	}

	r++
	ansi.SetCursor(r, c)
	fmt.Print("  " +
		ansi.CORNER_BOTTOM_LEFT +
		ansi.Repeat(21, ansi.LINE_HORIZONTAL) +
		ansi.CORNER_BOTTOM_RIGHT,
	)

	endRow, endCol = startRow + 13, startCol + 14
	return
}
