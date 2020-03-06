//Package section13

package section13

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func Dup1() {
	linesCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		linesCount[input.Text()]++
	}
	for key, val := range linesCount {
		fmt.Printf("%d\t%s\n", val, key)
	}
}
