package section13

import (
	"bufio"
	"fmt"
	"os"
)

func Dup2() {
	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) > 0 {
		for _, filePath := range files {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			CountLinesInFile(file, counts)
			file.Close()
		}
	} else {
		CountLinesInFile(os.Stdin, counts)
	}
	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, text)
		}
	}
}

func CountLinesInFile(fileReader *os.File, counts map[string]int) {
	input := bufio.NewScanner(fileReader)
	for input.Scan() {
		counts[input.Text()]++
	}
}
