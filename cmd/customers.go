/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/joukojo/go-what-did-i-do/services"
	"github.com/spf13/cobra"
)

// customersCmd represents the customers command
var customersCmd = &cobra.Command{
	Use:   "customers",
	Short: "Manage customers",
	Run: func(cmd *cobra.Command, args []string) {
		services.CustomerStorage.Print()
	},
}

func init() {
	rootCmd.AddCommand(customersCmd)

}
