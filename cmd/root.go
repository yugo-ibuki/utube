package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "utube",
	Short: "This is a simple CLI tool to handle youtube API commands.",
	Long:  "This is a simple CLI tool to handle youtube API commands.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("Hello, World!")
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
