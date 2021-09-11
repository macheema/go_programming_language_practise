package section12

import (
	"fmt"
	"os"
	"strings"
)

func Echo3() {
	sep := " "
	str := strings.Join(os.Args[1:], sep)
	fmt.Println(str)
}
