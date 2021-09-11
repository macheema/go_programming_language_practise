package test

import (
	"repeat/ch1/section12"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo3()
	}
}

func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		section12.Echo4()
	}
}
