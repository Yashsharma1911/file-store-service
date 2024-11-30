package cmd

import (
	"fmt"

	"github.com/Yashsharma1911/file-store-service/utils" // Import the utils package where MakeRequest is defined
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all files in the store",
	Run: func(cmd *cobra.Command, args []string) {
		respBody, err := utils.MakeRequest("GET", "http://localhost:8080/api/files", nil, "")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}

		fmt.Println(string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
