// Package cmd contains all command line options
package cmd

import (
	"strconv"
	"time"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// worksCmd represents the works command
var worksAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new work <task id> <description>",
	Run: func(cmd *cobra.Command, args []string) {
		services.WorkStorage.Print()

		if len(args) < 2 {
			cmd.Help()
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

		description := args[1]

		services.WorkStorage.Add(services.Work{
			ID:          time.Now().Unix(),
			TaskID:      taskID,
			Description: description,
			StartDate:   time.Now(),
		})

		_ = services.WorkStorage.SaveWorks()

	},
}

func init() {
	worksCmd.AddCommand(worksAddCmd)
}
