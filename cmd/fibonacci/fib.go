package main

import (
	"fmt"
	"os"
	"strconv"

	fibonacci "github.com/schnell18/play-golang/ch2/fib"
)

func main() {

	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseUint(arg, 10, 64); err == nil {
			fmt.Println(fibonacci.Fibonacci(int(n)))
		} else {
			fmt.Printf("bad input %s\n", arg)
		}
	}
}
