package cmd

import (
	"bytes"
	"fmt"
	"mime/multipart"
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

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		if err := utils.AddFileToWriter(writer, "file", filePath); err != nil {
			fmt.Println(err)
			return
		}

		_ = writer.Close()

		url := fmt.Sprintf("%s/api/files/%s", endpoint, fileName)
		respBody, err := utils.MakeRequest("PUT", url, body, writer.FormDataContentType())
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}

		fmt.Println(string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
