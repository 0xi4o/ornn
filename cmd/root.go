package cmd

import (
	"fmt"

	"github.com/0xi4o/ornn/log"
	"github.com/spf13/cobra"
)

var (
	logger log.Logger

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
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/ornn/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
