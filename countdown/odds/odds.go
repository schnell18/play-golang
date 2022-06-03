package main

import (
	"fmt"
)

func main() {
	// ch := make(chan int, 2) random result
	// ch := make(chan int) dead lock
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("%d ", x)
		case ch <- i:
		}
	}
}
