package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()
	filesizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, filesizes)
	}
	go func() {
		n.Wait()
		close(filesizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(files, bytes int64) {
	var sizeStr string
	if bytes < 1024 {
		sizeStr = fmt.Sprintf("%d B", bytes)
	} else if bytes < 1024*1024 {
		sizeStr = fmt.Sprintf("%.1f KB", float64(bytes/1024))
	} else if bytes < 1024*1024*1024 {
		sizeStr = fmt.Sprintf("%.1f MB", float64(bytes/1024/1024))
	} else if bytes < 1024*1024*1024*1024 {
		sizeStr = fmt.Sprintf("%.1f GB", float64(bytes/1024/1024/1024))
	}
	fmt.Printf("%d files  %s\n", files, sizeStr)
}

func walkDir(dir string, n *sync.WaitGroup, filesizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, filesizes)
		} else {
			filesizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
