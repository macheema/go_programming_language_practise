//Package section13 ...
package section13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Dup4 Modif y dup2 to print the names of all files in which each duplic ated line occ urs
func Dup4() {
	linesCount := make(map[string]int)
	linesinFiles := make(map[string]string)
	if len(os.Args) > 1 {
		for _, filePath := range os.Args[1:] {
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			countLinesWithFiles(file, filePath, linesCount, linesinFiles)
		}
	} else {
		countLines(os.Stdin, linesCount)
	}
	for key, val := range linesCount {
		if val > 1 {
			fmt.Printf("%s\t%d\t%s\n", linesinFiles[key], val, key)
		}
	}
}

func countLinesWithFiles(file *os.File, filePath string, linesCount map[string]int, linesinFiles map[string]string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		text := input.Text()
		linesCount[text]++
		if _, ok := linesinFiles[text]; ok {
			if exist := strings.Contains(linesinFiles[text], filePath); !exist {
				linesinFiles[text] += ", " + filePath
			}
		} else {
			linesinFiles[text] = filePath
		}
	}
}
