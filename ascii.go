package main

import (
	"errors"
	"strings"
)

const charHeight = 8
const linesPerChar = 9

func convertNewLines(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.ReplaceAll(input, "\\n", "\n")
}

func printAscii(input string, banner string) (string, error) {
	var result strings.Builder

	banner = strings.ReplaceAll(banner, "\r\n", "\n")
	bannerLines := strings.Split(banner, "\n")

	offset := 0
	if bannerLines[0] == "" {
		offset = 1
	}

	lastCharStart := 94*linesPerChar + offset
	if len(bannerLines) < lastCharStart+charHeight {
		return "", errors.New("banner file is incomplete")
	}

	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < charHeight; row++ {
			for i := 0; i < len(line); i++ {
				charNumber := int(line[i]) - 32
				start := charNumber*linesPerChar + offset

				result.WriteString(bannerLines[start+row])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
