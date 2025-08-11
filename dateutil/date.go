// Package dateutil provides utility functions for working with dates and times.
package dateutil

import (
	"fmt"
	"time"
)

// FormatDate formats a time.Time as a string in the format "YYYY-MM-DD".
func FormatDate(date *time.Time) string {
	// Format the date as "YYYY-MM-DD"
	if date == nil {
		return ""
	}
	return date.Format("2006-01-02")
}

// FormatTime formats a time.Time as a string in the format "YYYY-MM-DD HH:MM".
func FormatTime(date *time.Time) string {
	// Format the date as "YYYY-MM-DD"
	if date == nil {
		return ""
	}
	return date.Format("2006-01-02 15:04")
}

// HumanReadableDuration converts seconds to a human-readable string, e.g. 3661 -> "1h 1m 1s"
func HumanReadableDuration(seconds int64) string {
	if seconds < 0 {
		return ""
	}
	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60

	result := ""
	if h > 0 {
		result += fmt.Sprintf("%dh ", h)
	}
	if m > 0 || h > 0 {
		result += fmt.Sprintf("%dm ", m)
	}
	result += fmt.Sprintf("%ds", s)
	return result
}

// ParseDateTime parses a string as either a time (HH:MM) or a date+time (d.M.yyyy HH:MM).
// If only time is given, it uses today's date.
// Returns a time.Time and an error if parsing fails.
func ParseDateTime(input string) (time.Time, error) {
	// Try full date and time: d.M.yyyy HH:MM
	if t, err := time.Parse("2.1.2006 15:04", input); err == nil {
		return t, nil
	}

	// Try only time: HH:MM, use today
	today := time.Now()
	parsed, err := time.Parse("15:04", input)
	if err == nil {
		return time.Date(today.Year(), today.Month(), today.Day(), parsed.Hour(), parsed.Minute(), 0, 0, today.Location()), nil
	}

	return time.Time{}, fmt.Errorf("could not parse date/time: %q", input)
}
