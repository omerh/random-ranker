package cmd

import (
	"log"
	"random-ranker/helpers"

	"github.com/spf13/cobra"
)

var getRandom = &cobra.Command{
	Use:   "random",
	Short: "get random ranking",
	Run: func(cmd *cobra.Command, Args []string) {
		p1, _ := cmd.Flags().GetFloat64("p1")
		count, _ := cmd.Flags().GetInt("count")

		if !helpers.IsRequiredFlagsOK(p1) {
			return
		}

		log.Printf("Setting ranking for %v positions with P1 starting rank of %v", count, p1)
		helpers.RandomRanks(p1, count)
	},
}
