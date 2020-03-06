//Package section12 ...
package section12

import (
	"fmt"
	"os"
	"strings"
)

/*
Echo3 ...
*/
func Echo3() {
	var str, sep string = "", " "

	str = strings.Join(os.Args, sep)
	fmt.Println(str)
}
