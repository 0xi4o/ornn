package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all available project types",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\n  Available project types:\n\n")
			for _, projectType := range projectTypes {
				fmt.Printf("  %s\t%s\n", projectType.Command, projectType.Description)
			}
		},
	}

	projectTypes []ProjectType
)

func init() {
	projectTypes = GetProjectTypes()
}
