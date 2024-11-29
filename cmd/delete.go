package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/Yashsharma1911/file-store-service/utils"
)

var deleteCmd = &cobra.Command{
	Use:   "rm [file name]",
	Short: "Delete a file from the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		respBody, err := utils.MakeRequest("DELETE", fmt.Sprintf("http://localhost:8080/files/%s", fileName), nil, "")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}

		if string(respBody) == "File does not exist" {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("File deleted successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}