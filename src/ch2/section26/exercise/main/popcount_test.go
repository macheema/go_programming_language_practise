package main

import (
	"ch2/section26"
	"ch2/section26/exercise"
	"testing"
)

func BenchmarkPopCountExpression(b *testing.B) {
	// fmt.Println("BenchmarkPopCountExpression-> Value of BN=>", b.N)
	for i := 0; i < b.N; i++ {
		section26.PopCount(uint64(b.N))
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	// fmt.Println("BenchmarkPopCountWithLoop-> Value of BN=>", b.N)
	for i := 0; i < b.N; i++ {
		exercise.PopCountWithLoop(uint64(b.N))
	}
}

func BenchmarkPopCountRightShift(b *testing.B) {
	// fmt.Println("BenchmarkPopCountRightShift-> Value of BN=>", b.N)
	for i := 0; i < b.N; i++ {
		exercise.PopCountRightShift(uint64(b.N))
	}
}

func BenchmarkPopCountClearRightShift(b *testing.B) {
	// fmt.Println("BenchmarkPopCountClearRightShift-> Value of BN=>", b.N)
	for i := 0; i < b.N; i++ {
		exercise.PopCountClearRightShift(uint64(b.N))
	}
}
