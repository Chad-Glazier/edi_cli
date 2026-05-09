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

		for col := range 10 {
			s := ansi.FgBrightBlack(ansi.VACANT_SQUARE)

			pos := bb.Pos(row, col)
			isWhite := false
			isBlack := false
			for i := range 4 {
				if board.White[i] == pos {
					isWhite = true
					break
				}
				if board.Black[i] == pos {
					isBlack = true
					break
				}
			}

			switch {
			case isWhite:
				s = ansi.FgBrightCyan(ansi.WHITE_QUEEN)
				if board.Player == state.WHITE {
					s = ansi.Blink(s)
				}
			case isBlack:
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

	endRow, endCol = startRow+13, startCol+14
	return
}
