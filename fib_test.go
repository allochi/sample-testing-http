package main

import "testing"

func BenchmarkFibFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibFast(12)
	}
}

func BenchmarkFibRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRec(12)
	}
}
