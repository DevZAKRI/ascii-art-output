package main

import (
	"fmt"
	"os"
	"strings"

	output "output/artistTools"
)

// main is the entry point of the program.
// it validates the input argument, prints any validation errors,
// and then calls the Artist function to process the input.
func main() {
	// check if the input is valid
	valid, isoutput := output.ValidateArgument(os.Args)
	if valid != "" {
		fmt.Println(valid)
		return
	}
	template := "standard"
	input := ""
	fileName := ""
	if !isoutput {
		input = os.Args[1]
		if len(os.Args) == 3 && output.IsValidBanner(os.Args[2]) {
			if strings.HasSuffix(os.Args[2], ".txt") {
				template = os.Args[2][0 : len(os.Args[2])-4]
			} else {
				template = os.Args[2]
			}
		} else if len(os.Args) == 3 && !output.IsValidBanner(os.Args[2]) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	} else {
		input = os.Args[2]
		fileName = os.Args[1][9:]
		if input == "" {
			os.Create(fileName)
			return
		}
		if len(os.Args) == 4 && output.IsValidBanner(os.Args[3]) {
			if strings.HasSuffix(os.Args[3], ".txt") {
				template = os.Args[3][0 : len(os.Args[3])-4]
			} else {
				template = os.Args[3]
			}
		}
	}
	// take the input and and Process it
	result := output.Artist(input, template)
	// output the result either in console or in file
	output.OutputFinal(result, fileName)
}
