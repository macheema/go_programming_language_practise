//Package section13 ...
package section13

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Dup3 prints the count and text of lines that appear more than once
// in the input. It reads from a list of named files.
func Dup3() {
	linesCount := make(map[string]int)
	if len(os.Args) > 1 {
		for _, filePath := range os.Args[1:] {
			fileData, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			lines := strings.Split(string(fileData), "\n")
			for _, line := range lines {
				linesCount[line]++
			}
		}
	} else {
		countLines(os.Stdin, linesCount)
	}
	for key, val := range linesCount {
		fmt.Printf("%d\t%s\n", val, key)
	}
}
