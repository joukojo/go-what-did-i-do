// Package cmd contains all command line options
package cmd

import (
	"fmt"
	"time"

	"github.com/joukojo/go-what-did-i-do/promptutil"
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
	Run: func(cmd *cobra.Command, _ []string) {
		// Import the Customer type from the appropriate package

		if customerName == "" {
			fmt.Println("Please enter the customer's name:")
			customerName = promptutil.AskString("Name: ")
		}
		if customerName == "" {
			fmt.Println("Error: Customer name cannot be empty.")
			_ = cmd.Help()
			return
		}
		services.CustomerStorage.Add(services.Customer{
			ID:   time.Now().Unix(),
			Name: customerName,
		})
		err := services.CustomerStorage.Save()
		if err != nil {
			fmt.Println("Error saving customer:", err)
		}
	},
}

func init() {
	customersCmd.AddCommand(customersAddCmd)
	customersAddCmd.Flags().StringVar(&customerName, "name", "", "Name of the customer")
}
