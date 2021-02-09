package main

import (
	"fmt"
	"github.com/alex19861108/Tiger/cmd"
	"os"

	"github.com/spf13/cobra"
)

var versionFlag bool

var RootCmd = &cobra.Command{
	Use: "Tiger",
	RunE: func(c *cobra.Command, args []string) error {
		if versionFlag {
			fmt.Printf("%s version %s, build %s\n", BinaryName, Version, GitCommit)
			return nil
		}
		return c.Usage()
	},
}

func init() {
	RootCmd.AddCommand(cmd.Start)

	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "show version")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
