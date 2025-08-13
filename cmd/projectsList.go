// Package cmd contains all command line options
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(_ *cobra.Command, _ []string) {
		services.ProjectStorage.PrintProjects()
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
