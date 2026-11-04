package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Just one input is needed")
		return
	}

	input := os.Args[1]

	// Read the banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	banner := string(data)

	// Turn the \n from the command into real newlines
	input = convertNewLines(input)

	// Print the input
	printAscii(input, banner)
}
