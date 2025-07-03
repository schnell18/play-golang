package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	defer fmt.Println("solar system")
	defer fmt.Println("galaxy")
	defer fmt.Println("universe")
	fmt.Println("Hello")
	deferLIFODemo()
}

func deferLIFODemo() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
