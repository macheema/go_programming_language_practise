package section12

import (
	"fmt"
	"os"
)

func Echo4() {
	fmt.Println(os.Args[1:])
}
