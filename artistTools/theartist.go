package output

import (
	"fmt"
	"os"
	"strings"
)

// Artist is a function that generates ASCII art based on the given input, template, and file name.
// It reads the ASCII art template from a file, replaces any carriage return characters, and splits the input into lines.
// If the input string is empty after removing "\\n", it removes the first line.
// It then save each line as ASCII art using the LineAsAscii function in variable finalAsciiArt
// and return it to main
func Artist(input string, template string) string {
	// check for errors too in case file no longer exist
	asciiGraph, err := ReadFile("banners/" + template + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	for idx, LINE := range asciiGraph {
		asciiGraph[idx] = strings.ReplaceAll(LINE, "\r", "")
	}
	lines := strings.Split(input, "\\n")
	// If the input string is empty after removing "\\n", it removes the first line.
	if strings.ReplaceAll(input, "\\n", "") == "" {
		lines = lines[1:]
	}
	finalAsciiArt := ""
	for _, line := range lines {
		// add each line to finalAsciiArt using the LineAsAscii function.
		finalAsciiArt += LineAsAscii(line, asciiGraph)
	}
	return finalAsciiArt
}
