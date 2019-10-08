package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {

    for _, arg := range os.Args[1:] {
        if n, err := strconv.ParseUint(arg, 10, 64); err == nil {
            fmt.Println(Fib(n))
        } else {
            fmt.Printf("bad input %s\n", arg)
        }
    }
}

func Fib(n uint64) uint64 {
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
