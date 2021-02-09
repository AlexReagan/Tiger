package cmd

import (
	"github.com/spf13/cobra"
)

var Start = &cobra.Command{
	Use:   "start",
	Short: "Start Tiger modules",
	Long: `
Start Tiger and run until a stop command is received.
	`,
	RunE:          start,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func start(c *cobra.Command, args []string) error {
	return nil
}

