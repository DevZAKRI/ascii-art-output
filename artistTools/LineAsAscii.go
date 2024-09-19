package output

// LineAsAscii prints a given line of text as ASCII art using a provided ASCII graph.
// It takes the line of text, the ASCII graph, and an optional filename as input.
// The function converts each character in the line to its corresponding ASCII representation
// and appends it to the final ASCII art. The ASCII art is then returned to artist
//
// Parameters:
//   - line: The line of text to be converted to ASCII art.
//   - asciiGraph: The ASCII graph used for converting characters to ASCII art.

func LineAsAscii(line string, asciiGraph []string) string {
	var asciiChars []string
	finalAsciiArt := ""
	if line != "" {
		for _, char := range line {
			// add each chars to the asciiChars line by line
			for i := 8; i > 0; i-- {
				asciiChars = append(asciiChars, string(asciiGraph[findLastLine(char)-i]))
			}
		}

		for i := 0; i < 8; i++ {
			for j := 0; j < len(asciiChars); j += 8 {
				finalAsciiArt += asciiChars[i+j]
			}
			finalAsciiArt += "\n"
		}
	} else {
		finalAsciiArt += "\n"
	}
	return finalAsciiArt
}

// find the last line after the char
func findLastLine(char rune) int {
	return int((char - 31) * (9))
}
