// Package cmd contains all command line options
package cmd

import (
	"fmt"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

var customerID int64

// customersDelCmd represents the customersDel command
var customersDelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a customer",

	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Deleting customer #", customerID)
		services.CustomerStorage.Delete(customerID)
		err := services.CustomerStorage.Save()
		if err != nil {
			fmt.Println("Error saving customer after deleting customer:", err)
		}
	},
}

func init() {
	customersDelCmd.Flags().Int64Var(&customerID, "id", -1, "ID of the customer to delete")
	if err := customersDelCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println("Error marking flag as required:", err)
	}
	customersCmd.AddCommand(customersDelCmd)
}
