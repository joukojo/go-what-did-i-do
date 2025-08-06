/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

var (
	customerName string
)

// customersAddCmd represents the customersAdd command
var customersAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new customer",
	Long: `This command requires the following parameters: 
	- name: The name of the customer`,
	Run: func(cmd *cobra.Command, args []string) {
		// Import the Customer type from the appropriate package
		services.CustomerStorage.Add(services.Customer{
			ID:   time.Now().Unix(),
			Name: customerName,
		})
		services.CustomerStorage.Save()
	},
}

func init() {
	customersCmd.AddCommand(customersAddCmd)
	customersAddCmd.Flags().StringVar(&customerName, "name", "", "Name of the project")
}
