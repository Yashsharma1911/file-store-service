package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/Yashsharma1911/file-store-service/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file paths]",
	Short: "Add files to the store",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for _, filePath := range args {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", filePath, err)
				continue
			}
			defer file.Close()

			part, err := writer.CreateFormFile("files", filePath)
			if err != nil {
				fmt.Printf("Error creating form file for %s: %v\n", filePath, err)
				continue
			}

			_, err = io.Copy(part, file)
			if err != nil {
				fmt.Printf("Error copying content for %s: %v\n", filePath, err)
				continue
			}
		}

		_ = writer.Close()
		url := fmt.Sprintf("%s/api/files", endpoint)

		respBody, err := utils.MakeRequest("POST", url, body, writer.FormDataContentType())
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		fmt.Println(string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
