// Package cmd contains all command line options
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// taskListCmd represents the tasks command
var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(_ *cobra.Command, _ []string) {
		services.TaskStorage.PrintTasks()

	},
}

func init() {
	tasksCmd.AddCommand(taskListCmd)

}
