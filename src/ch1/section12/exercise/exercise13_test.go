//Package exercise ...
package exercise

import (
	"ch1/section12"
	"testing"
)

/*BenchmarkEcho1
C:\Users\awais\AppData\Local\Temp\go-build036253158\b001\exercise.test.exe -test.benchmem=true -test.run=^$ -test.bench=^(BenchmarkEcho1)$
  100000	     16382 ns/op	     368 B/op	       4 allocs/op
PASS
Success: Benchmarks passed.
*/
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo1()
	}
}

/*BenchmarkEcho2
C:\Users\awais\AppData\Local\Temp\go-build278055690\b001\exercise.test.exe -test.benchmem=true -test.run=^$ -test.bench=^(BenchmarkEcho2)$
  100000	     18054 ns/op	     368 B/op	       4 allocs/op
PASS
ok  	ch1/exercise	2.045s
Success: Benchmarks passed.
*/
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo2()
	}
}

/*BenchmarkEcho3
C:\Users\awais\AppData\Local\Temp\go-build086539950\b001\exercise.test.exe -test.benchmem=true -test.run=^$ -test.bench=^(BenchmarkEcho3)$
  100000	     17212 ns/op	     160 B/op	       2 allocs/op
PASS
ok  	ch1/exercise	1.984s
Success: Benchmarks passed.
*/
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo3()
	}
}

/*BenchmarkEcho4
[-test.benchmem=true -test.run=^$ -test.bench=^(BenchmarkEcho4)$]
  200000	     16199 ns/op	      80 B/op	       4 allocs/op
PASS
ok  	ch1/exercise	3.402s
Success: Benchmarks passed.
*/
func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo4()
	}
}
