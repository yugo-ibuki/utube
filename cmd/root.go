package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"utube/channel"
	"utube/youtube"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "utube",
	Short: "This is a simple CLI tool to handle youtube API commands.",
	Long:  "This is a simple CLI tool to handle youtube API commands.",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		y := youtube.New()

		chInfo := make([]channel.Member, len(youtube.ChannelIds))

		for i, chId := range youtube.ChannelIds {
			response, err := y.GetChCall(chId)
			if err != nil {
				log.Fatalf("Error making API call: %v", err)
			}
			chInfo[i] = channel.Member{
				Id:           response.Items[i].Id,
				Title:        response.Items[i].Snippet.Title,
				ChannelId:    response.Items[i].Id,
				ThumbnailUrl: response.Items[i].Snippet.Thumbnails.Default.Url,
				Views:        response.Items[i].Statistics.ViewCount,
			}
		}

		for _, c := range chInfo {
			fmt.Printf("This channel's ID is %s. Its title is '%s', and it has %v views.\n",
				c.Id, c.Title, c.Views)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// フラグ
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
