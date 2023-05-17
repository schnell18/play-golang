package main

import (
	"io"
	"log"
	"os"

	"github.com/schnell18/play-golang/ch13/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("zipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("zipper: close %v\n", err)
	}
}
