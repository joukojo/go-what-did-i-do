// Package cmd contains all command line options
package cmd

import (
	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage tasks",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

}
