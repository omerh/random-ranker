package cmd

import "github.com/spf13/cobra"

var (
	getExample = `
	# Get Random Ranks
	random-ranker get random
	`
	getShort = ("random-ranker getRandom")
)

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   getShort,
	Example: getExample,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
		}
	},
}

func init() {
	getCmd.AddCommand(getRandom)
}
