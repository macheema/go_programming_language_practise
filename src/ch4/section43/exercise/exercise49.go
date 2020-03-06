package exercise

import (
	"bufio"
	"fmt"
	"os"
)

/*Exercise49 Write a program word freq to rep ort the frequency of each word in an input text
file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
words instead of lines.
*/
func Exercise49() {

	wordCount := map[string]int{}
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for {
		if read := in.Scan(); read {
			word := in.Text()
			wordCount[word]++
		} else {
			break
		}
	}
	fmt.Printf("word\tcount\n")
	for c, n := range wordCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
