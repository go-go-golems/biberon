package main

import (
	"embed"
	"github.com/go-go-golems/biberon/cmd/biberon/cmds"
	"github.com/go-go-golems/glazed/pkg/help"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "biberon",
	Short: "biberon is a tool to manage bibliography data",
}

func main() {
	_ = rootCmd.Execute()
}

//go:embed doc/*
var docFS embed.FS

func init() {
	helpSystem := help.NewHelpSystem()
	err := helpSystem.LoadSectionsFromFS(docFS, ".")
	if err != nil {
		panic(err)
	}

	helpSystem.SetupCobraRootCommand(rootCmd)

	rootCmd.AddCommand(cmds.BibtexCmd)
}
