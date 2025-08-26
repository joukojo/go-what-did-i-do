// Package cmd contains all command line options
package cmd

import (
	"strconv"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// worksCmd represents the works command
var worksStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop work <work id> <description>",
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

		if len(args) > 1 {
			description := args[1]
			services.WorkStorage.Stop(workID, description)
		} else {
			services.WorkStorage.Stop(workID, "")
		}

		_ = services.WorkStorage.SaveWorks()

	},
}

func init() {
	worksCmd.AddCommand(worksStopCmd)
}
