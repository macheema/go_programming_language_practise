//Package section12 ...
package section12

import (
	"fmt"
	"os"
)

/*
Echo1 ...
*/
func Echo1() {
	var str, sep string
	for i := 0; i < len(os.Args); i++ {
		str += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(str)
}
