package main

import (
	"strings"
)

func convertNewLines(input string) string {
	return strings.ReplaceAll(input, "\\n", "\n")
}

func printAscii(input string, banner string) string {
	var result strings.Builder

	bannerLines := splitLines(banner)
	inputLines := splitLines(input)

	for _, line := range inputLines {

		// Empty line
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Print the 8 rows of ASCII characters
		for row := 0; row < 8; row++ {

			for i := 0; i < len(line); i++ {
				char := line[i]

				charNumber := int(char) - 32
				start := charNumber*9 + 1

				result.WriteString(bannerLines[start+row])
			}

			result.WriteString("\n")
		}
	}

	return result.String()
}

func splitLines(text string) []string {
	lines := []string{}
	current := ""

	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(text[i])
		}
	}

	lines = append(lines, current)

	return lines
}
