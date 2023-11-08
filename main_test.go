package main

import "testing"

func Benchmark_f1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2(numOfGorutines)
	}
}
