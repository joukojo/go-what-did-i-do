// Package cmd contains all command line options
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version bool = false

// These variables can be set at build time using -ldflags
var (
	GitCommit string = "none"
	GitBranch string = "local"
	BuildTime string = "unknown"
)

func PrintVersion() {
	fmt.Println("What did I do?")

	fmt.Printf("Commit: %s Branch: %s Build Time: %s\n", GitCommit, GitBranch, BuildTime)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "what-did-i-do",
	Short: "A simple cli tool for time tracking ",
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			PrintVersion()
			return
		}
	},
}

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
