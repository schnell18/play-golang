package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte

	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()

	fmt.Println("Echo from main")
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
