package exercise

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hashFlag *string = flag.String("hash", "256", "generates hash of StdIn")

//Exercise42 ...
func Exercise42() {
	fmt.Printf("Please enter text to generates a hash of it : ")
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	flag.Parse()
	switch *hashFlag {
	case "384":
		fmt.Printf("sha384 Hash : %v\n", sha512.Sum384([]byte(line)))
	case "512":
		fmt.Printf("sha512 Hash : %v", sha512.Sum512([]byte(line)))
	default:
		fmt.Printf("sha256 Hash : %v", sha256.Sum256([]byte(line)))
	}
}
