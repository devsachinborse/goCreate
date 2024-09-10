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
		fmt.Printf("Creating project: %s\n", projectName)
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
		Options: []string{"Standard", "Flat", "Clean Architecture", "Layered", "Domain-Driven Design(DDD)","hexagonal", "microservices"},
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
	green := color.New(color.FgGreen).SprintFunc()   // Success color
	red := color.New(color.FgRed).SprintFunc()       // Error color
	yellow := color.New(color.FgYellow).SprintFunc() // Folder creation message

	// Define folder structures based on architecture
	var dirs []string
	var filesContent map[string]string

	switch architecture {
	case "Standard":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "pkg"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Standard architecture")
}
`,
			filepath.Join(projectName, "pkg", "example.go"): `
package pkg

// Example package
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a standard project structure.
`,
		}
	case "Flat":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "services"),
			filepath.Join(projectName, "utils"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Flat architecture")
}
`,
			filepath.Join(projectName, "services", "service.go"): `
package services

// Service implementation
`,
			filepath.Join(projectName, "utils", "utils.go"): `
package utils

// Utility functions
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a flat project structure.
`,
		}
	case "Layered":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "app"),
			filepath.Join(projectName, "internal", "domain"),
			filepath.Join(projectName, "internal", "repository"),
			filepath.Join(projectName, "internal", "service"),
			filepath.Join(projectName, "internal", "handler"),
			filepath.Join(projectName, "internal", "utils"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Layered architecture")
}
`,
			filepath.Join(projectName, "internal", "app", "app.go"): `
package app

// Application logic
`,
			filepath.Join(projectName, "internal", "domain", "domain.go"): `
package domain

// Domain models
`,
			filepath.Join(projectName, "internal", "repository", "repository.go"): `
package repository

// Repository implementation
`,
			filepath.Join(projectName, "internal", "service", "service.go"): `
package service

// Service implementation
`,
			filepath.Join(projectName, "internal", "handler", "handler.go"): `
package handler

// HTTP handlers
`,
			filepath.Join(projectName, "internal", "utils", "utils.go"): `
package utils

// Utility functions
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a layered project structure.
`,
		}
	case "Domain-Driven Design(DDD)":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "domains"),
			filepath.Join(projectName, "services"),
			filepath.Join(projectName, "handlers"),
			filepath.Join(projectName, "repositories"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("DDD architecture")
}
`,
			filepath.Join(projectName, "domains", "domain.go"): `
package domains

// Domain models
`,
			filepath.Join(projectName, "services", "service.go"): `
package services

// Service implementation
`,
			filepath.Join(projectName, "handlers", "handler.go"): `
package handlers

// HTTP handlers
`,
			filepath.Join(projectName, "repositories", "repository.go"): `
package repositories

// Repository implementation
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a DDD project structure.
`,
		}
	case "Clean Architecture":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "controller"),
			filepath.Join(projectName, "internal", "service"),
			filepath.Join(projectName, "internal", "repository"),
			filepath.Join(projectName, "internal", "model"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Clean architecture")
}
`,
			filepath.Join(projectName, "internal", "controller", "controller.go"): `
package controller

// Controller logic
`,
			filepath.Join(projectName, "internal", "service", "service.go"): `
package service

// Service logic
`,
			filepath.Join(projectName, "internal", "repository", "repository.go"): `
package repository

// Repository logic
`,
			filepath.Join(projectName, "internal", "model", "model.go"): `
package model

// Domain models
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a Clean architecture project structure.
`,
		}

	case "microservices":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "services", "service1"),
			filepath.Join(projectName, "services", "service2"),
			filepath.Join(projectName, "services", "service3"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Microservices architecture")
}
`,
			filepath.Join(projectName, "services", "service1", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Service 1")
}
`,
			filepath.Join(projectName, "services", "service2", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Service 2")
}
`,
			filepath.Join(projectName, "services", "service3", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Service 3")
}
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a Microservices architecture project structure.
`,
		}
	case "hexagonal":
		dirs = []string{
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "adapter", "inbound"),
			filepath.Join(projectName, "internal", "adapter", "outbound"),
			filepath.Join(projectName, "internal", "application"),
			filepath.Join(projectName, "internal", "domain"),
		}
		filesContent = map[string]string{
			filepath.Join(projectName, "cmd", "main.go"): `
package main

import "fmt"

func main() {
	fmt.Println("Hexagonal architecture")
}
`,
			filepath.Join(projectName, "internal", "adapter", "inbound", "http.go"): `
package inbound

// HTTP inbound adapter
`,
			filepath.Join(projectName, "internal", "adapter", "outbound", "repository.go"): `
package outbound

// Repository outbound adapter
`,
			filepath.Join(projectName, "internal", "application", "service.go"): `
package application

// Application services
`,
			filepath.Join(projectName, "internal", "domain", "model.go"): `
package domain

// Domain models
`,
			filepath.Join(projectName, "go.mod"): `
module ` + projectName + `

go 1.20
`,
			filepath.Join(projectName, "README.md"): `
# ` + projectName + `

This is a Hexagonal architecture project structure.
`,
		}
	default:
		fmt.Println("Unknown architecture option. Use one of: standard, flat, layered, ddd, clean")
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

	// // Create some common files (e.g., README, go.mod, etc.)
	// files := []string{
	// 	filepath.Join(projectName, "README.md"),
	// 	filepath.Join(projectName, "go.mod"),
	// }

	for file, content := range filesContent {
		f, err := os.Create(file)
		if err != nil {
			fmt.Printf("%s %s: %v\n", red("Error creating file"), file, err)
			return
		}
		_, err = f.WriteString(content)
		if err != nil {
			fmt.Printf(red("Error wrting to file %s: %v\n"), file, err)
			return
		}
		f.Close()
		fmt.Println(yellow("Created file:"), file)
	}

	fmt.Printf("%s Project structure for %s architecture created successfully!\n", green("Success!"), green(architecture))
}
