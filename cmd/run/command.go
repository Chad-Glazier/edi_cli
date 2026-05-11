package run

import (
	"errors"
	"fmt"
	"time"

	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi/state"
	"github.com/Chad-Glazier/edi_cli/cmd/flags"
	"github.com/Chad-Glazier/edi_cli/out"
	"github.com/Chad-Glazier/edi_cli/sim"
	"github.com/Chad-Glazier/edi_cli/ui"
	"github.com/spf13/cobra"
)

func RunCommand() *cobra.Command {

	var turnTimer time.Duration
	var whiteName, blackName flags.VIName

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start a game between two programs",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if whiteName != "" && blackName == "" {
				return errors.New("no VI was specified for Black; use --black")
			}
			if blackName != "" && whiteName == "" {
				return errors.New("no VI was specified for White; use --white")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			var white, black edi.VI
			if whiteName == "" && blackName == "" {
				white, black = ui.SelectPlayers()
			} else {
				white = flags.CreateVI(whiteName)
				black = flags.CreateVI(blackName)
			}

			width, height := ui.TerminalSize()
			ui.ClearScreen()
			ui.HideCursor()
			defer ui.ShowCursor()

			boardCh := sim.Game(white, black, turnTimer)
			var activePlayer state.PlayerColor
			for board := range boardCh {
				out.PrintState(board, height/2-7, width/2-12)
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

	cmd.Flags().DurationVarP(
		&turnTimer, "time", "t", 5*time.Second, "time limit per turn")
	cmd.Flags().VarP(
		&whiteName, "white", "w", "a VI player: edi, arrow, or random")
	cmd.Flags().VarP(
		&blackName, "black", "b", "a VI player: edi, arrow, or random")

	return cmd
}
