package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/Yashsharma1911/file-store-service/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [file path]",
	Short: "Update a file in the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		fileName := filepath.Base(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", fileName)
		if err != nil {
			fmt.Printf("Error creating form file: %v\n", err)
			return
		}

		_, err = io.Copy(part, file)
		if err != nil {
			fmt.Printf("Error copying file content: %v\n", err)
			return
		}

		_ = writer.Close()

		url := fmt.Sprintf("http://localhost:8080/files/%s", fileName)
		respBody, err := utils.MakeRequest("PUT", url, body, writer.FormDataContentType())
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}

		fmt.Println("Response:", string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
