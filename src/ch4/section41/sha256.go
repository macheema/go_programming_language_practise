package section41

import (
	"crypto/sha256"
	"fmt"
)

//SHA256 ...
func SHA256(code1 string, code2 string) (hash1 [32]uint8, hash2 [32]uint8) {
	c1 := sha256.Sum256([]byte(code1))
	c2 := sha256.Sum256([]byte(code2))
	fmt.Printf("c1 -> %x, %v, %t, %T\n", c1, c1, c1 == c2, c1)
	fmt.Printf("c2 -> %x, %v, %t, %T\n", c2, c2, c1 == c2, c2)
	hash1 = c1
	hash2 = c2
	return
}
