package run

import (
	"errors"
	"fmt"
	"time"

	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi/state"
	"github.com/Chad-Glazier/edi_cli/ansi"
	"github.com/Chad-Glazier/edi_cli/out"
	"github.com/Chad-Glazier/edi_cli/sim"
	"github.com/Chad-Glazier/edi_cli/ui"
	"github.com/spf13/cobra"
)

type VIName string

const (
	EDI    VIName = "edi"
	ARROW  VIName = "arrow"
	RANDOM VIName = "random"
)

func (v *VIName) String() string {
	return string(*v)
}

func (v *VIName) Set(s string) error {
	switch s {
	case "edi", "arrow", "random":
		*v = VIName(s)
		return nil
	default:
		return fmt.Errorf(`must be one of "edi", "arrow", or "random"`)
	}
}

func (v *VIName) Type() string {
	return "VI"
}

func RunCommand() *cobra.Command {

	var turnTimer time.Duration
	var whiteName, blackName VIName

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
				switch whiteName {
				case EDI:
					white = &edi.EDI{}
				case ARROW:
					white = &edi.Arrow{}
				case RANDOM:
					white = &edi.Random{}
				}
				switch blackName {
				case EDI:
					black = &edi.EDI{}
				case ARROW:
					black = &edi.Arrow{}
				case RANDOM:
					black = &edi.Random{}
				}
			}

			width, height := ui.TerminalSize()
			ansi.ClearScreen()
			ansi.HideCursor()
			defer ansi.ShowCursor()

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
