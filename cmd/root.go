package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/utube/channel"
	"github.com/yugo-ibuki/utube/youtube"
	"log"
	"os"
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

		chInfo := make([]channel.Channel, len(youtube.ChannelIds))

		for i, chId := range youtube.ChannelIds {
			response, err := y.DoChCall(chId)
			if err != nil {
				log.Fatalf("Error making API call: %v", err)
			}
			chInfo[i] = channel.Format(i, response)
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
