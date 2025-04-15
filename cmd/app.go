package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Model struct {
	projectName string
	err         error
}

var (
	appCmd = &cobra.Command{
		Use:   "app",
		Short: "Create a new fullstack application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating new fullstack application...")
		},
	}
)

func initializeModel(projectName string) (*Model, error) {
	model := &Model{
		projectName: projectName,
		err:         nil,
	}
	return model, nil
}

func init() {}
