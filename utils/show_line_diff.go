package utils

import (
	"fmt"
	"strings"
)

// ShowLineDiff prints a color-coded, line-by-line diff between two multi-line strings.
// It splits both inputs on '\n' and compares corresponding lines; when lines differ it
// prints the old line (removal) in red prefixed with " - " and the new line (addition)
// in green prefixed with " + ". Line numbers are 1-based. Identical lines are omitted.
// The function writes directly to stdout using ANSI color codes and does not return any value.
func ShowLineDiff(old string, new string) {
	oldLines := strings.Split(old, "\n")
	newLines := strings.Split(new, "\n")

	maxLen := max(len(oldLines), len(newLines))

	for i := 0; i < maxLen; i++ {
		var oldLine, newLine string
		if i < len(oldLines) {
			oldLine = oldLines[i]
		}
		if i < len(newLines) {
			newLine = newLines[i]
		}

		if oldLine != newLine {
			if oldLine != "" {
				fmt.Printf("\033[31m [%d] - %s\033[0m\n", i+1, oldLine)
			}
			if newLine != "" {
				fmt.Printf("\033[32m [%d] + %s\033[0m\n\n", i+1, newLine)
			}
		}
	}
}
