package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	// Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	// ValidArgs: []string{"app", "cli", "server"},
	Run: func(cmd *cobra.Command, args []string) {
		runCreate()
	},
}

type CreateModel struct {
	choices  map[int]Choice
	cursor   int
	selected int
}

type Choice struct {
	label string
	value string
}

func initCreateModel() CreateModel {
	return CreateModel{
		choices: map[int]Choice{
			1: {label: "API server", value: "api"},
			2: {label: "Fullstack Application", value: "app"},
			3: {label: "CLI Application", value: "cli"},
		},
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
			_, ok := createModel.choices[createModel.cursor]
			if ok {
				createModel.selected = createModel.cursor
			} else {
				createModel.cursor = 0
			}
		case tea.KeyEnter.String():
			if createModel.selected != 0 {
				fmt.Printf("Selected option: %v", createModel.choices[createModel.selected])
				appCmd.Execute()
			}
		}
	}
	return createModel, nil
}

func (createModel CreateModel) View() string {
	view := "\nWhat type of project are you starting?\n\n"

	for i, choice := range createModel.choices {
		cursor := " "
		if createModel.cursor == i {
			cursor = ">"
		}

		checked := " "
		if createModel.selected != 0 {
			checked = "x"
		}

		view += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.label)
	}

	view += "\nPress Ctrl-C or q to quit.\n"

	return view
}

func runCreate() {
	p := tea.NewProgram(initCreateModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

func init() {
	createCmd.AddCommand(appCmd)
	createCmd.AddCommand(cliCmd)
	createCmd.AddCommand(serverCmd)
}
