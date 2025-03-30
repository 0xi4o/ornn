package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:       "create",
	Short:     "Create a new project",
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"app", "cli", "server"},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Creating new project...")
	// },
}

func init() {
	createCmd.AddCommand(appCmd)
	createCmd.AddCommand(cliCmd)
	createCmd.AddCommand(serverCmd)
}
