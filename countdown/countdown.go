package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	// waiting user abort
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for cnt := 10; cnt > 0; cnt-- {
		fmt.Println(cnt)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	fmt.Println("Lauching...")
}
