package section43

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

//CharacterCount ...
func CharacterCount() {
	invalidChar := 0
	count := map[rune]int{}
	var utflen [utf8.UTFMax + 1]int

	in := bufio.NewReader(os.Stdin)
	for {
		r, nOfBytes, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "character count :\t%s\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && nOfBytes == 1 {
			invalidChar++
			continue
		}
		count[r]++
		utflen[nOfBytes]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range count {
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
