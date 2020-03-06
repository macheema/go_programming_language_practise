package exercise

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

//Exercise48 Modif y charcount to count letters, digits, and so on in their Unico de categories,using functions like unicode.IsLetter
// Control Characters
// Digit Characters
// Letter Characters
// Mark Characters
// Number Characters all type numbers (e.g. Telgu, Arabic .etc)
// Punctuation characters
// Space Character
// Symbol Character
//
func Exercise48() {

	characterCount := map[rune]int{}
	categoryCount := map[string]int{}
	var utflen [unicode.MaxRune + 1]int
	in := bufio.NewReader(os.Stdin)
	invalidChar := 0
	for {
		r, n, er := in.ReadRune()
		if er == io.EOF {
			break
		}
		if er != nil {
			fmt.Fprintf(os.Stderr, "Error\t:\t%v\n", er)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalidChar++
			continue
		}
		for key, cat := range unicode.Categories {
			if unicode.Is(cat, r) {
				categoryCount[key]++
				characterCount[r]++
				utflen[n]++
				break
			}
		}
	}
	fmt.Printf("category\tcount\n")
	for c, n := range categoryCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range characterCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalidChar > 0 {
		fmt.Printf("\n%d invalid UTF8characters\n", invalidChar)
	}
}
