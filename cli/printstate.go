package cli

import (
	"fmt"

	"github.com/Chad-Glazier/edi/bb"
	"github.com/Chad-Glazier/edi/state"
)

func PrintState(board *state.Board) {
	fmt.Println()
	fmt.Println("\t    0 1 2 3 4 5 6 7 8 9 ")
	fmt.Println("\t  " +
		CORNER_TOP_LEFT +
		repeat(21, LINE_HORIZONTAL) +
		CORNER_TOP_RIGHT,
	)

	for row := range 10 {
		fmt.Printf("\t%d "+LINE_VERTICAL, row)
		for col := range 10 {
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
		fmt.Print(" " + LINE_VERTICAL + "\n")
	}

	fmt.Println("\t  " +
		CORNER_BOTTOM_LEFT +
		repeat(21, LINE_HORIZONTAL) +
		CORNER_BOTTOM_RIGHT,
	)
	fmt.Println()
}
