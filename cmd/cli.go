package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Create a new CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new CLI application...")
	},
}

func init() {
}
