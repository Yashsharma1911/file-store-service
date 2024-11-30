package cmd

import (
	"fmt"

	"github.com/Yashsharma1911/file-store-service/utils"
	"github.com/spf13/cobra"
)

var wordCountCmd = &cobra.Command{
	Use:   "wc",
	Short: "Get word count of all files present in file store",
	Run: func(cmd *cobra.Command, args []string) {
		respBody, err := utils.MakeRequest("GET", "http://localhost:8080/wc", nil, "")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}

		fmt.Println(string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(wordCountCmd)
}
