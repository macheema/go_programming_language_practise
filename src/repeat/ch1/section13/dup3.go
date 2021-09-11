package section13

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Dup3() {
	files := os.Args[1:]
	counts := make(map[string]int)
	for _, filePath := range files {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}

	for line, num := range counts {
		fmt.Printf("%d\t%s\n", num, line)
	}
}
