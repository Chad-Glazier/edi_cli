package out

import (
	"fmt"

	"github.com/Chad-Glazier/edi/bb"
	"github.com/Chad-Glazier/edi/state"
)

func PrintState(board *state.Board, startRow, startCol int) {
	r := startRow
	c := startCol

	// Header row
	setCursor(r, c)
	fmt.Print("    0 1 2 3 4 5 6 7 8 9 ")

	// Top border
	r++
	setCursor(r, c)
	fmt.Print("  " +
		CORNER_TOP_LEFT +
		repeat(21, LINE_HORIZONTAL) +
		CORNER_TOP_RIGHT,
	)

	// Board rows
	for row := 0; row < 10; row++ {
		r++
		setCursor(r, c)

		fmt.Printf("%d %s", row, LINE_VERTICAL)

		for col := 0; col < 10; col++ {
			s := fgBrightBlack("\u00B7")

			switch {
			case board.White.Flagged(bb.Pos(row, col)):
				s = fgBrightCyan("\u25A0")
				if board.Player == state.WHITE {
					s = blink(s)
				}
			case board.Black.Flagged(bb.Pos(row, col)):
				s = fgBrightRed("\u25A0")
				if board.Player == state.BLACK {
					s = blink(s)
				}
			case board.Occupancy.Flagged(bb.Pos(row, col)):
				s = fgWhite("\u2715")
			}

			fmt.Print(" " + s)
		}

		fmt.Print(" " + LINE_VERTICAL)
	}

	// Bottom border
	r++
	setCursor(r, c)
	fmt.Print("  " +
		CORNER_BOTTOM_LEFT +
		repeat(21, LINE_HORIZONTAL) +
		CORNER_BOTTOM_RIGHT,
	)
}
