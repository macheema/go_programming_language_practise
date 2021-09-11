package exercise

import (
	"fmt"
	"repeat/ch1/section12"
	"time"
)

func Exercise13() {
	startTime := time.Now()
	section12.Echo1()
	fmt.Printf("%v nanoseconds spent during Echo1", time.Since(startTime))
	startTime = time.Now()
	section12.Echo2()
	fmt.Printf("%v nanoseconds spent during Echo2", time.Since(startTime))
	startTime = time.Now()
	section12.Echo3()
	fmt.Printf("%v nanoseconds spent during Echo3", time.Since(startTime))
	startTime = time.Now()
	section12.Echo4()
	fmt.Printf("%v nanoseconds spent during Echo4\n", time.Since(startTime))

}
