package run

import (
	"fmt"
	"time"

	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi/state"
	"github.com/Chad-Glazier/edi_cli/ansi"
	"github.com/Chad-Glazier/edi_cli/out"
	"github.com/Chad-Glazier/edi_cli/sim"
	"github.com/spf13/cobra"
)

func RunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start a game between two programs",
		RunE: func(cmd *cobra.Command, args []string) error {
			white := &edi.EDI{}
			black := &edi.Random{}
			turnTimer := time.Second * 10

			ansi.ClearScreen()

			boardCh := sim.Game(white, black, turnTimer)
			var activePlayer state.PlayerColor
			for board := range boardCh {
				out.PrintState(board, 2, 4)
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
