//Package section13 ...
package section13

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
func Dup2() {
	linesCount := make(map[string]int)
	if len(os.Args) > 1 {
		for _, filePath := range os.Args[1:] {
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			countLines(file, linesCount)
		}
	} else {
		countLines(os.Stdin, linesCount)
	}
	for key, val := range linesCount {
		if val > 1 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}

func countLines(file *os.File, linesCount map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		linesCount[input.Text()]++
	}
}
