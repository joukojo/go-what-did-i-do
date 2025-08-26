// Package cmd contains all command line options
package cmd

import (
	"strconv"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// worksCmd represents the works command
var worksDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete work <work id>",
	Run: func(cmd *cobra.Command, args []string) {
		workID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			cmd.PrintErrln("Invalid WorkID, must be an integer")
			return
		}

		if !services.WorkStorage.Exists(workID) {
			cmd.PrintErrln("Work with ID", workID, "does not exist.")
			return
		}

		services.WorkStorage.Remove(workID)

		_ = services.WorkStorage.SaveWorks()

	},
}

func init() {
	worksCmd.AddCommand(worksDeleteCmd)
}
