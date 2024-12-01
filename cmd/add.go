package cmd

import (
	"bytes"
	"fmt"
	"mime/multipart"

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
			if err := utils.AddFileToWriter(writer, "files", filePath); err != nil {
				fmt.Println(err)
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
