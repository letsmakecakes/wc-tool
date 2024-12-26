package utils

import (
	"fmt"
	"os"
	"strings"
)

// CheckError handles error checking and exit if error occurs
func CheckError(err error, message string) {
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		if err != nil {
			fmt.Printf("Write error encountered: %v", err)
		}
		os.Exit(1)
	}
}

// FormatOutput formats the counts and filename according to wc standard
func FormatOutput(counts []int, filename string) string {
	output := ""
	for _, count := range counts {
		output += fmt.Sprintf("%d ", count)
	}
	output = strings.TrimSpace(output) // Remove trailing space
	if filename != "" {
		output += fmt.Sprintf(" %s", filename)
	}
	return output + "\n"
}
