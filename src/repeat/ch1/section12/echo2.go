package section12

import (
	"fmt"
	"os"
)

func Echo2() {
	var str, sep string = "", ""
	for _, v := range os.Args[1:] {
		str += sep + v
		sep = " "
	}
	fmt.Println(str)
}
