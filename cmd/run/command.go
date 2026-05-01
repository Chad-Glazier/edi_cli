package run

import (
	"fmt"
	"os"
	"time"

	"github.com/Chad-Glazier/edi"
	"github.com/Chad-Glazier/edi_cli/output"
	"github.com/spf13/cobra"
)

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a game between EDI and herself",
		Run: func(cmd *cobra.Command, args []string) {

			timeLimit, err := cmd.Flags().GetUint("time")
			if err != nil {
				fmt.Println("Invalid time value.")
				os.Exit(1)
			}

			output.RunGame(
				&edi.EDI{},
				&edi.EDI{},
				time.Duration(timeLimit) * time.Second,
			)
		},
	}
	cmd.Flags().Uint("time", 5, "the time limit per turn in seconds")
	return cmd
}
