package cmd

import (
	"strconv"
	"time"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// tasksAddCmd represents the tasks add command
var tasksAddCmd = &cobra.Command{
	Use:   "add <name> <description> <projectID>",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for adding a new task
		if len(args) < 3 {
			_ = cmd.Help()
			return
		}

		projectID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			cmd.PrintErrln("Invalid ProjectID, must be an integer")
			return
		}

		if !services.ProjectStorage.Exists(projectID) {
			cmd.PrintErrln("Project with ID", projectID, "does not exist.")
			return
		}

		task := services.Task{
			ID:          time.Now().Unix(),
			Name:        args[0],
			Description: args[1],
			ProjectID:   projectID,
		}
		services.TaskStorage.Add(task)
		_ = services.TaskStorage.SaveTasks()
	},
}

func init() {
	tasksCmd.AddCommand(tasksAddCmd)

}
