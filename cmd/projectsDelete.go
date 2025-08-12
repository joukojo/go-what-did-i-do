// Package cmd contains all command line options
package cmd

import (
	"strconv"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// projectsDeleteCmd represents the projects delete command
var projectsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for deleting a project
		if len(args) < 1 {
			_ = cmd.Help()
			return
		}

		projectID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			cmd.PrintErrln("Invalid ProjectID, must be an integer")
			return
		}

		if !services.ProjectStorage.Exists(projectID) {
			cmd.PrintErrln("Project with ID", projectID, "does not exist.")
			return
		}

		services.ProjectStorage.Delete(projectID)
		_ = services.ProjectStorage.SaveProjects()
	},
}

func init() {
	projectsCmd.AddCommand(projectsDeleteCmd)

}
