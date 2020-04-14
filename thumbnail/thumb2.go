package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func main() {
	total := makeThumbnails2(os.Args[1:])
	fmt.Printf("Total thumbnail file size: %d\n", total)
}

func makeThumbnails2(filenames []string) int64 {

	sizes := make(chan int64)
	var wg sync.WaitGroup

	for _, f := range filenames {
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
