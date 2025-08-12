// Package cmd contains all command line options
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// customersCmd represents the customers command
var customersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List customers",
	Run: func(_ *cobra.Command, _ []string) {
		services.CustomerStorage.Print()
	},
}

func init() {
	customersCmd.AddCommand(customersListCmd)
}
