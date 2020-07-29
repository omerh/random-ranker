package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCommand = `
		# Get Random
		random-ranker list regions regions		
	`
)

var rootCmd = &cobra.Command{
	Use:     "random-ranker",
	Short:   "random-ranker for setting ranks",
	Example: rootCommand,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// Persistent flags goes here
func init() {
	// Commands
	rootCmd.AddCommand(getCmd)

	// Flags
	rootCmd.PersistentFlags().Float64("p1", 0, "P1 Rank")
	rootCmd.MarkFlagRequired("p1")
	rootCmd.PersistentFlags().Int("count", 20, "How many positions to rank")
	rootCmd.MarkFlagRequired("count")
}

// Execute using cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
