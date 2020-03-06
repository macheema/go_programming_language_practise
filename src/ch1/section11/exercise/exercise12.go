//Package exercise ...
package exercise

import (
	"fmt"
	"os"
)

//Exercise1_2 ...
func Exercise1_2() {
	for index, value := range os.Args {
		fmt.Printf(" [%v] : %v\n", index, value)
	}
}
