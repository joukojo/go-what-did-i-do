package dateutil_test

import (
	"testing"
	"time"

	"github.com/joukojo/go-what-did-i-do/dateutil"
)

func TestFormatTime(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected string
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: "",
		},
		{
			name: "valid date and time",
			input: func() *time.Time {
				tm := time.Date(2024, 6, 1, 14, 30, 0, 0, time.UTC)
				return &tm
			}(),
			expected: "2024-06-01 14:30",
		},
		{
			name: "midnight time",
			input: func() *time.Time {
				tm := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
				return &tm
			}(),
			expected: "2023-12-31 00:00",
		},
		{
			name: "single digit hour and minute",
			input: func() *time.Time {
				tm := time.Date(2022, 1, 2, 3, 4, 0, 0, time.UTC)
				return &tm
			}(),
			expected: "2022-01-02 03:04",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dateutil.FormatTime(tt.input)
			if result != tt.expected {
				t.Errorf("FormatTime(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestHumanReadableDuration tests the HumanReadableDuration function for various cases.
func TestHumanReadableDuration(t *testing.T) {
	tests := []struct {
		seconds  int64
		expected string
	}{
		{0, "0s"},
		{1, "1s"},
		{59, "59s"},
		{60, "1m 0s"},
		{61, "1m 1s"},
		{3599, "59m 59s"},
		{3600, "1h 0m 0s"},
		{3661, "1h 1m 1s"},
		{7322, "2h 2m 2s"},
		{-1, ""}, // negative seconds
	}

	for _, tt := range tests {
		result := dateutil.HumanReadableDuration(tt.seconds)
		if result != tt.expected {
			t.Errorf("HumanReadableDuration(%d) = %q, want %q", tt.seconds, result, tt.expected)
		}
	}
}

// TestParseDateTime tests the ParseDateTime function for various input formats.
func TestParseDateTime(t *testing.T) {
	today := time.Now()
	todayDate := time.Date(today.Year(), today.Month(), today.Day(), 15, 15, 0, 0, today.Location())

	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "only time, today",
			input:    "15:15",
			expected: todayDate,
			wantErr:  false,
		},
		{
			name:     "full date and time",
			input:    "8.5.2025 15:15",
			expected: time.Date(2025, 5, 8, 15, 15, 0, 0, today.Location()),
			wantErr:  false,
		},
		{
			name:    "invalid input",
			input:   "not a date",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := dateutil.ParseDateTime(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseDateTime(%q) expected error, got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Errorf("ParseDateTime(%q) unexpected error: %v", tt.input, err)
				return
			}
			if result.Year() != tt.expected.Year() || result.Month() != tt.expected.Month() || result.Day() != tt.expected.Day() || result.Hour() != tt.expected.Hour() || result.Minute() != tt.expected.Minute() {
				t.Errorf("ParseDateTime(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
