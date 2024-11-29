package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "store",
	Short: "CLI for interacting with the file store service",
	Long:  "A command-line interface for adding, listing, updating, and deleting files in the file store service",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}