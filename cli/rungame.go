package cli

import (
	"fmt"
	"time"

	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi/state"
)

func RunGame(white, black edi.VI, turnTimer time.Duration) {
	board := state.InitialState()
	clearScreen()
	PrintState(&board)

	player := state.WHITE

	for len(board.Successors()) != 0 {
		var move *state.Move
		if player == state.WHITE {
			move = white.Consult(&board, turnTimer)
			player = state.BLACK
		} else {
			move = black.Consult(&board, turnTimer)
			player = state.WHITE
		}
		board.Apply(move)
		clearScreen()
		PrintState(&board)
	}

	if player == state.WHITE {
		fmt.Println("Black Wins")
	} else {
		fmt.Println("White Wins")
	}
}
