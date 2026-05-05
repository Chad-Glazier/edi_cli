package sim

import (
	"time"

	"github.com/Chad-Glazier/edi/diag"
	"github.com/Chad-Glazier/edi/state"
)

// Makes a search algorithm play against itself, collecting diagnostic data
// on each move. The first returned channel tracks the board state, while the
// second receives the reports for each search.
func DiagGame[WhiteReport, BlackReport diag.Report](
	white diag.SearchFunc[WhiteReport],
	black diag.SearchFunc[BlackReport],
	turnTimer time.Duration,
) (<-chan state.Board, <-chan WhiteReport, <-chan BlackReport) {

	boardCh := make(chan state.Board)
	whiteReportCh := make(chan WhiteReport)
	blackReportCh := make(chan BlackReport)

	go func() {
		defer close(boardCh)
		defer close(whiteReportCh)
		defer close(blackReportCh)

		board := state.InitialState()
		boardCh <- board

		player := board.Player
		for len(board.Successors()) != 0 {

			var move state.Move
			
			if player == state.WHITE {
				report := white(&board, turnTimer)
				move = report.Move()
				player = state.BLACK
				whiteReportCh <- report

			} else {
				report := black(&board, turnTimer)
				move = report.Move()
				player = state.WHITE
				blackReportCh <- report
			}

			board.Apply(&move)
			boardCh <- board
		}
	}()

	return boardCh, whiteReportCh, blackReportCh
}
