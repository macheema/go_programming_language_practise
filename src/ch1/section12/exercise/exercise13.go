//Package exercise ...
package exercise

import (
	"ch1/section12"
	"fmt"
	"time"
)

//Exercise1_3 ...
func Exercise1_3() {
	start := time.Now()
	section12.Echo1()
	fmt.Printf("%v elapsed\n", time.Since(start).Nanoseconds())
	//
	start = time.Now()
	section12.Echo2()
	fmt.Printf("%v elapsed\n", time.Since(start).Nanoseconds())
	//
	start = time.Now()
	section12.Echo3()
	fmt.Printf("%v elapsed\n", time.Since(start).Nanoseconds())
	//
	start = time.Now()
	section12.Echo4()
	fmt.Printf("%v elapsed\n", time.Since(start).Nanoseconds())
}
