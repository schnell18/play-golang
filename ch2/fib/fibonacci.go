package fib

func Fibonacci(n uint64) uint64 {
	var i, x, y uint64 = 0, 0, 1
	for i = 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func Fibonacci2(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
