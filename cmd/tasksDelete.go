package cmd

import (
	"strconv"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// tasksDeleteCmd represents the tasks delete command
var tasksDeleteCmd = &cobra.Command{
	Use:   "delete <taskID>",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for deleting a task
		if len(args) < 1 {
			_ = cmd.Help()
			return
		}

		taskID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			cmd.PrintErrln("Invalid TaskID, must be an integer")
			return
		}

		if !services.TaskStorage.Exists(taskID) {
			cmd.PrintErrln("Task with ID", taskID, "does not exist.")
			return
		}

		services.TaskStorage.Delete(taskID)
		_ = services.TaskStorage.SaveTasks()
	},
}

func init() {
	tasksCmd.AddCommand(tasksDeleteCmd)

}
