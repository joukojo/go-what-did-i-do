package cmd

import (
	"strconv"
	"time"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// projectsAddCmd represents the projects add command
var projectsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new project",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for adding a new project
		if len(args) < 3 {
			_ = cmd.Help()
			return
		}

		customerID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			cmd.PrintErrln("Invalid CustomerID, must be an integer")
			return
		}

		if !services.CustomerStorage.Exists(customerID) {
			cmd.PrintErrln("Customer with ID", customerID, "does not exist.")
			return
		}

		project := services.Project{
			ID:          time.Now().Unix(),
			Name:        args[0],
			Description: args[1],
			CustomerID:  customerID,
		}
		services.ProjectStorage.Add(project)
		_ = services.ProjectStorage.SaveProjects()
	},
}

func init() {
	projectsCmd.AddCommand(projectsAddCmd)

}
