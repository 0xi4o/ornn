package main

import (
	"fmt"
	"os"

	"github.com/0xi4o/ornn/log"
	"github.com/spf13/cobra"
)

type ProjectType struct {
	command     string
	description string
}

var (
	logger       log.Logger
	projectTypes []ProjectType
	version      = "0.1.0"

	rootCmd = &cobra.Command{
		Use:   "ornn",
		Short: "Ornn is a fullstack web application framework for Go",
		// TODO: add a long description
		// Long: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger = log.Initialize()
			defer logger.Close()

			if err := cmd.Usage(); err != nil {
				return fmt.Errorf("failed to print usage: %v", err)
			}
			return nil
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all available project types",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\n  Available project types:\n\n")
			for _, projectType := range projectTypes {
				fmt.Printf("  %s\t%s\n", projectType.command, projectType.description)
			}
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ornn version: %s\n", version)
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	projectTypes = []ProjectType{
		{command: "api", description: "REST API server"},
		{command: "app", description: "Fullstack Application with Go backend and React frontend"},
		{command: "cli", description: "CLI application powered by Cobra and Bubbletea"},
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/ornn/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(versionCmd)
}
