package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Yashsharma1911/file-store-service/server/models"
	"github.com/Yashsharma1911/file-store-service/utils"
	"github.com/spf13/cobra"
)

var frequencyCountCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "Get word count of all files present in file store",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the limit and order flags
		limit, _ := cmd.Flags().GetString("limit")
		order, _ := cmd.Flags().GetString("order")

		// Prepare the query parameters
		if limit == "" {
			limit = "10"
		}
		if order == "" {
			order = "dsc"
		}

		respBody, err := utils.MakeRequest("GET", fmt.Sprintf("%s/api/frequent?limit=%s&order=%s", endpoint, limit, order), nil, "")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		var wordCount []models.Word
		_ = json.Unmarshal(respBody, &wordCount)

		if len(wordCount) == 0 {
			fmt.Println("No files found.")
			return
		}
		for _, word := range wordCount {
			// Convert size to KB for better readability
			count := word.Count
			fmt.Printf("%s\t%d\n", word.Word, count)
		}
	},
}

func init() {
	rootCmd.AddCommand(frequencyCountCmd)

	// Add flags for limit and order
	frequencyCountCmd.PersistentFlags().String("order", "dsc", "Order of word frequency (asc or dsc)")
	frequencyCountCmd.PersistentFlags().StringP("limit", "n", "10", "Number of most/least frequent words to retrieve")
}
