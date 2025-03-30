package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Create a new fullstack application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new fullstack application...")
	},
}

func init() {}
