// Package cmd contains all command line options
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// customersCmd represents the customers command
var customersCmd = &cobra.Command{
	Use:   "customers",
	Short: "Manage customers",
	Run: func(cmd *cobra.Command, _ []string) {
		services.CustomerStorage.Print()
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(customersCmd)

}
