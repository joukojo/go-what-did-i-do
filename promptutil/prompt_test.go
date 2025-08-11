package promptutil_test

import (
	"os"
	"testing"

	"github.com/joukojo/go-what-did-i-do/promptutil"
)

func TestAskString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "User writes text",
			input:    "John\n",
			expected: "John",
		},
		{
			name:     "User presses enter without any text",
			input:    "\n",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original Stdin
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			// Replace Stdin with our mock input using os.Pipe
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}
			_, err = w.Write([]byte(tt.input))
			if err != nil {
				t.Fatalf("failed to write to pipe: %v", err)
			}
			_ = w.Close()
			os.Stdin = r

			// Run the function
			got := promptutil.AskString("What is your name?")

			// Compare results
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
