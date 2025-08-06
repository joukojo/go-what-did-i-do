package main

import (
	"fmt"
)

var (
	GitCommit string
	GitBranch string
	BuildTime string
)

func PrintVersion() {
	fmt.Printf("Commit: %s\nBranch: %s\nBuild Time: %s\n", GitCommit, GitBranch, BuildTime)
}

func main() {
	fmt.Println("What did I do?")
	PrintVersion()

}
