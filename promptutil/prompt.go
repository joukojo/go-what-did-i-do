// Package promptutil provides utility functions for prompting user input.
package promptutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AskString prints a question and returns the user's input (trimmed).
func AskString(question string) string {
	fmt.Print(question + " ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
