// Package cmd contains all command line options
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = false

// These variables can be set at build time using -ldflags
var (
	GitCommit = "none"
	GitBranch = "local"
	BuildTime = "unknown"
)

// PrintVersion prints the version information of the application.
func PrintVersion() {
	fmt.Println("What did I do?")

	fmt.Printf("Commit: %s Branch: %s Build Time: %s\n", GitCommit, GitBranch, BuildTime)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "what-did-i-do",
	Short: "A simple cli tool for time tracking ",
	Run: func(_ *cobra.Command, _ []string) {
		if version {
			PrintVersion()
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVar(&version, "version", false, "Print version information")
}
