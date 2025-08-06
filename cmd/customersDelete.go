/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

var customerId int64

// customersDelCmd represents the customersDel command
var customersDelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a customer",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting customer #", customerId)
		services.CustomerStorage.Delete(customerId)
		services.CustomerStorage.Save()
	},
}

func init() {
	customersDelCmd.Flags().Int64Var(&customerId, "id", -1, "ID of the customer to delete")
	customersDelCmd.MarkFlagRequired("id")
	customersCmd.AddCommand(customersDelCmd)

}
