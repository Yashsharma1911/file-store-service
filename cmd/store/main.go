package main

import (
	"os"

	"github.com/Yashsharma1911/file-store-service/cmd"
)

// main is the entrypoint of the file store command-line tool
func main() {
	// Execute the root command
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
