package cmd

import (
	"github.com/spf13/cobra"
)

var worksCmd = &cobra.Command{
	Use:   "works",
	Short: "Manage Works",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(worksCmd)
}
