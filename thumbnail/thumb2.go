package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func main() {
	files := make(chan string)
	for _, f := range os.Args[1:] {
		files <- f
	}
	total := makeThumbnails2(files)
	fmt.Printf("Total thumbnail file size: %d", total)
}

func makeThumbnails2(filenames <-chan string) int64 {

	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
