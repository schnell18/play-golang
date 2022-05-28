package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type link struct {
	depth int
	url   string
}

type pack struct {
	depth int
	urls  []string
}

// add depth control
func main() {
	maxDepth := 2
	worklist := make(chan pack)
	unseenLinks := make(chan link)

	go func() { worklist <- pack{depth: 1, urls: os.Args[1:]} }()

	for i := 0; i < 20; i++ {
		go func() {
			for lk := range unseenLinks {
				var pk pack
				if lk.depth <= maxDepth {
					pk = crawl(lk)
				}
				go func() { worklist <- pk }()
			}
		}()
	}

	seen := make(map[string]bool)
	npacks := 1
	for list := range worklist {
		for _, url := range list.urls {
			if !seen[url] {
				seen[url] = true
				npacks++
				unseenLinks <- link{depth: list.depth, url: url}
			}
		}
		npacks--
		if npacks == 0 {
			close(worklist)
			close(unseenLinks)
		}
	}
}

func crawl(l link) pack {
	fmt.Printf("[%d]: %s\n", l.depth, l.url)
	list, err := links.Extract(l.url)
	if err != nil {
		log.Print(err)
	}
	return pack{depth: l.depth + 1, urls: list}
}
