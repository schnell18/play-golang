package main

import (
	"fmt"

	"github.com/schnell18/play-golang/cgo/print"
	"github.com/schnell18/play-golang/cgo/random"
)

func main() {
	random.Seed(1000)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", random.Random())
	}

	print.Print("Good is good!")
}
