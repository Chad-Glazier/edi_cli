package run

import (
	"log"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi_cli/cmd/flags"
	"github.com/spf13/cobra"
)

func RunCommand() *cobra.Command {

	var turnTimer time.Duration
	var whiteName, blackName flags.VIName

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start a game between two programs",
		RunE: func(cmd *cobra.Command, args []string) error {

			var white, black edi.VI
			if whiteName != "" {
				white = flags.CreateVI(whiteName)
			}
			if blackName != "" {
				black = flags.CreateVI(blackName)
			}

			p := tea.NewProgram(NewGameModel(white, black, &turnTimer))
			_, err := p.Run()
			if err != nil {
				log.Fatal(err.Error())
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
