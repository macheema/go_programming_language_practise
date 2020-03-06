//Package section12 ...
package section12

import (
	"fmt"
	"os"
)

/*
Echo2 ...
*/
func Echo2() {
	var str, sep string

	for _, arg := range os.Args {
		str += sep + arg
		sep = " "
	}
	fmt.Println(str)
}
