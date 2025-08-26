// Package cmd contains all command line options
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// worksCmd represents the works command
var worksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List works",
	Run: func(_ *cobra.Command, _ []string) {
		services.WorkStorage.Print()
	},
}

func init() {
	worksCmd.AddCommand(worksListCmd)
}
