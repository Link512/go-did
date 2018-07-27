package cmd

import (
	"github.com/Link512/godid"
	"github.com/spf13/cobra"
)

var lastWeekCmd = &cobra.Command{
	Use:   "lastWeek",
	Short: "Displays the tasks logged last week",
	RunE: func(cmd *cobra.Command, args []string) error {
		flat, err := cmd.Flags().GetBool("flat")
		if err != nil {
			return err
		}
		lastWeek, err := godid.GetLastWeek(flat)
		if err != nil {
			return err
		}
		printResults(lastWeek)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lastWeekCmd)
	lastWeekCmd.Flags().BoolP("flat", "f", false, "Do not aggregate the tasks per day")
}
