package run

import (
	"fmt"
	"time"

	"github.com/Chad-Glazier/edi/state"
	"github.com/Chad-Glazier/edi_cli/ansi"
	"github.com/Chad-Glazier/edi_cli/out"
	"github.com/Chad-Glazier/edi_cli/sim"
	"github.com/Chad-Glazier/edi_cli/ui"
	"github.com/spf13/cobra"
)

func RunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start a game between two programs",
		RunE: func(cmd *cobra.Command, args []string) error {

			white, black := ui.SelectPlayers()

			turnTimer := time.Second * 10

			ansi.ClearScreen()

			boardCh := sim.Game(white, black, turnTimer)
			var activePlayer state.PlayerColor
			termWidth, termHeight := ui.TerminalSize()
			for board := range boardCh {
				out.PrintState(board, termHeight/2-7, termWidth/2-12)
				activePlayer = board.Player
			}

			if activePlayer == state.WHITE {
				fmt.Println("\nBlack Wins!")
			} else {
				fmt.Println("\nWhite Wins!")
			}

			return nil
		},
	}
	return cmd
}
