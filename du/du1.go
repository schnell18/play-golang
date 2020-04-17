package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkDir(dir string, filesizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, filesizes)
		} else {
			filesizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"/Users/zhangfeng775/research"}
	}

	filesizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, filesizes)
		}
		close(filesizes)
	}()

	var nfiles, nbytes int64
	for size := range filesizes {
		nfiles++
		nbytes += size
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
