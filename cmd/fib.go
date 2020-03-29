package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/schnell18/play-golang/fibonacci"
)

func main() {

	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseUint(arg, 10, 64); err == nil {
			fmt.Println(fibonacci.Fibonacci(n))
		} else {
			fmt.Printf("bad input %s\n", arg)
		}
	}
}
