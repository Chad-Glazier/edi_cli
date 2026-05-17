package analyze

import (
	"log"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/Chad-Glazier/edi_cli/cmd/flags"
	"github.com/spf13/cobra"
)

func AnalyzeCommand() *cobra.Command {

	var turnTimer time.Duration
	var vi flags.VI

	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Have a program play against itself to collect analytics",
		RunE: func(cmd *cobra.Command, args []string) error {

			p := tea.NewProgram(
				NewGameModel(vi, turnTimer),
			)
			_, err := p.Run()
			if err != nil {
				log.Fatal(err.Error())
			}

			return nil
		},
	}

	cmd.Flags().DurationVarP(
		&turnTimer, "time", "t", 0, "time limit per turn")
	cmd.Flags().VarP(
		&vi, "vi", "v", flags.VI_USAGE)

	return cmd
}
