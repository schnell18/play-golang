package fib

import (
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	cases := []struct {
		arg  uint64
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2},
		{4, 3}, {5, 5}, {6, 8}, {7, 13},
		{8, 21}, {9, 34}, {10, 55}, {11, 89},
		{12, 144}, {13, 233}, {14, 377},
	}
	for j := 0; j < b.N; j++ {
		for _, tc := range cases {
			_ = Fibonacci(tc.arg)
		}
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	cases := []struct {
		arg  uint64
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2},
		{4, 3}, {5, 5}, {6, 8}, {7, 13},
		{8, 21}, {9, 34}, {10, 55}, {11, 89},
		{12, 144}, {13, 233}, {14, 377},
	}
	for j := 0; j < b.N; j++ {
		for _, tc := range cases {
			_ = Fibonacci2(tc.arg)
		}
	}
}
