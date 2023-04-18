package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		arg  uint64
		want uint64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
	}
	for i, tc := range cases {
		assert.Equal(tc.want, Fibonacci(tc.arg))
	}
}

func FuzzTestFibonacci(f *testing.F) {
	testcases := []int{1, 2, 4, 5, 10, 13}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, n int) {
		assert := assert.New(t)
		ret := Fibonacci(n)
		if n == 0 {
			assert.Equal(0, ret)
		} else if n == 1 || n == 2 {
			assert.Equal(1, ret)
		} else {
			assert.Equal(ret, Fibonacci(n-1)+Fibonacci(n-2))
		}
	})
}
