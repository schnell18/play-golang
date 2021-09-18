// Package fibonacci calculates fibonacci numbers
package fibonacci

// Fibonacci Calculate nth fibonacci number
func Fibonacci(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	var a uint64 = 1
	var b uint64 = 1
	var i uint64 = 3
	for i = 3; i <= n; i = i + 1 {
		a, b = b, a+b
	}
	return b
}
