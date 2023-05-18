package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			// value := value
			ch <- &value
		}
		// close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) // what will be printed here?
		}
		wg.Done()
	}()

	wg.Wait()
}
