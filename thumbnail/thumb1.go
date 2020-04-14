package main

import (
	"fmt"
	"os"

	"gopl.io/ch8/thumbnail"
)

func main() {
	files, err := makeThumbnails(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to create thumbnail due to %v", err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Printf("generated thumbnail: %s\n", file)
	}
}

func makeThumbnails(filenames []string) ([]string, error) {

	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	var thumbnails []string
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnails = append(thumbnails, it.thumbfile)
	}
	return thumbnails, nil
}
