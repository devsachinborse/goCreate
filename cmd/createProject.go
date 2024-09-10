package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the create-project command
var createProjectCmd = &cobra.Command{
	Use:   "create-project [project name]",
	Short: "Create a new Golang project with a predefined structure",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is passed (project name)
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		architecture := promptArchitectureSelection()
		createProjectStructure(projectName, architecture)
	},
}

func init() {
	rootCmd.AddCommand(createProjectCmd)
}

// promptArchitectureSelection prompts the user to select an architecture type
func promptArchitectureSelection() string {
	var architecture string

	// Customize prompt message color
	cyan := color.New(color.FgCyan).SprintFunc()

	prompt := &survey.Select{
		Message: cyan("Select the project architecture:"),
		Options: []string{"Standard Layout", "Flat Structure", "Layered Structure", "Domain-Driven Design", "Clean Architecture"},
	}
	err := survey.AskOne(prompt, &architecture)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red("Error selecting architecture:"), err)
		os.Exit(1)
	}
	return architecture
}

// createProjectStructure creates the project folder structure based on the selected architecture
func createProjectStructure(projectName, architecture string) {
	green := color.New(color.FgGreen).SprintFunc()  // Success color
	red := color.New(color.FgRed).SprintFunc()      // Error color
	yellow := color.New(color.FgYellow).SprintFunc() // Folder creation message

	// Define folder structures based on architecture
	var dirs []string

	switch architecture {
	case "Standard Layout":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "app"),
			filepath.Join(projectName, "internal", "domain"),
			filepath.Join(projectName, "internal", "repository"),
			filepath.Join(projectName, "internal", "service"),
			filepath.Join(projectName, "internal", "handler"),
			filepath.Join(projectName, "internal", "utils"),
			filepath.Join(projectName, "pkg"),
		}
	case "Flat Structure":
		dirs = []string{
			filepath.Join(projectName),
		}
	case "Layered Structure":
		dirs = []string{
			filepath.Join(projectName, "app"),
			filepath.Join(projectName, "domain"),
			filepath.Join(projectName, "service"),
			filepath.Join(projectName, "repository"),
		}
	case "Domain-Driven Design":
		dirs = []string{
			filepath.Join(projectName, "domain", "model"),
			filepath.Join(projectName, "domain", "service"),
			filepath.Join(projectName, "infrastructure", "repository"),
			filepath.Join(projectName, "application", "usecases"),
			filepath.Join(projectName, "interfaces", "api"),
		}
	case "Clean Architecture":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "entities"),
			filepath.Join(projectName, "internal", "usecases"),
			filepath.Join(projectName, "internal", "interfaces"),
			filepath.Join(projectName, "internal", "infrastructure"),
			filepath.Join(projectName, "internal", "controller"),
		}
	default:
		fmt.Println(red("Unknown architecture. Exiting."))
		return
	}

	// Create directories
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("%s %s: %v\n", red("Error creating directory"), dir, err)
			return
		}
		fmt.Println(yellow("Created directory:"), dir)
	}

	// Create some common files (e.g., README, go.mod, etc.)
	files := []string{
		filepath.Join(projectName, "README.md"),
		filepath.Join(projectName, "go.mod"),
	}

	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			fmt.Printf("%s %s: %v\n", red("Error creating file"), file, err)
			return
		}
		f.Close()
		fmt.Println(yellow("Created file:"), file)
	}

	fmt.Printf("%s Project structure for %s architecture created successfully!\n", green("Success!"), green(architecture))
}
