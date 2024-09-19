package output

import (
	"fmt"
	"io"
	"os"
)

// OutputFinal print our FInal Ascii Graph either in console or in file
func OutputFinal(result string, fileName string) {
	if fileName != "" {
		file, _ := os.Create(fileName)
		io.WriteString(file, result)
	} else {
		fmt.Print(result)
	}
}
