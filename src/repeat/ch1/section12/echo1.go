package section12

import (
	"fmt"
	"os"
)

func Echo1() {
	var str, sep string
	for i := 1; i < len(os.Args); i++ {
		str += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(str)
}
