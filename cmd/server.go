package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a new server application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new server application...")
	},
}

func init() {
}
