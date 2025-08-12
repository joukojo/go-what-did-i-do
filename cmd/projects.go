// Package cmd contains all command line options
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage projects",
	Run: func(cmd *cobra.Command, _ []string) {
		services.ProjectStorage.PrintProjects()
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)

}
