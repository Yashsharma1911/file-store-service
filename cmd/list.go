package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Yashsharma1911/file-store-service/server/models"
	"github.com/Yashsharma1911/file-store-service/utils" // Import the utils package where MakeRequest is defined
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all files in the store",
	Run: func(cmd *cobra.Command, args []string) {
		respBody, err := utils.MakeRequest("GET", fmt.Sprintf("%s/api/files", endpoint), nil, "")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		var fileMetadata []models.FileMetadata
		_ = json.Unmarshal(respBody, &fileMetadata)

		if len(fileMetadata) == 0 {
			fmt.Println("No files found.")
			return
		}
		for _, file := range fileMetadata {
			size := file.Size
			fmt.Printf("%s\t%s\n\n", "File Name", "File Size")
			fmt.Printf("%s\t%d\n", file.FileName, size)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
