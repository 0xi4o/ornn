package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type ProjectType struct {
	Command     string
	Description string
}

type CreateModel struct {
	choices  []ProjectType
	cursor   int
	selected int
}

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		// Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		// ValidArgs: []string{"app", "cli", "server"},
		Run: func(cmd *cobra.Command, args []string) {
			// p := tea.NewProgram(model)
			// if _, err := p.Run(); err != nil {
			// 	fmt.Printf("Error running program: %v", err)
			// 	os.Exit(1)
			// }
			cmd.Usage()
		},
	}
	model CreateModel
)

func GetProjectTypes() []ProjectType {
	return []ProjectType{
		{Command: "api", Description: "REST API server"},
		{Command: "app", Description: "Fullstack Application with Go backend and React frontend"},
		{Command: "cli", Description: "CLI application powered by Cobra and Bubbletea"},
	}
}

func initCreateModel() CreateModel {
	return CreateModel{
		choices:  GetProjectTypes(),
		cursor:   0,
		selected: 0,
	}
}

func (createModel CreateModel) Init() tea.Cmd {
	return nil
}

func (createModel CreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), "q":
			return createModel, tea.Quit
		case tea.KeyUp.String(), "k":
			if createModel.cursor > 1 {
				createModel.cursor--
			}
		case tea.KeyDown.String(), "j":
			if createModel.cursor < 3 {
				createModel.cursor++
			}
		case tea.KeySpace.String():
			projectType := createModel.choices[createModel.cursor]
			if projectType.Command != "" && projectType.Description != "" {
				createModel.selected = createModel.cursor
			} else {
				createModel.cursor = 0
			}
		case tea.KeyEnter.String():
			fmt.Printf("Selected option: %v", createModel.choices[createModel.selected])
			appCmd.Execute()
		}
	}
	return createModel, nil
}

func (createModel CreateModel) View() string {
	view := "\nWhat type of project are you starting?\n\n"

	for i, projectType := range createModel.choices {
		cursor := " "
		if createModel.cursor == i {
			cursor = ">"
		}

		checked := " "
		if createModel.selected != 0 {
			checked = "x"
		}

		view += fmt.Sprintf("%s [%s] %s\n", cursor, checked, projectType.Command)
	}

	view += "\nPress Ctrl-C or q to quit.\n"

	return view
}

func init() {
	// model = initCreateModel()
	createCmd.AddCommand(appCmd)
}
