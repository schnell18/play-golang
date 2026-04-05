// Explore the experimental goroutine leak profile in 1.26
package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/schnell18/play-golang/goroutineleak"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	for {
		var ws []goroutineleak.WorkItem
		for i := 0; i < 100; i++ {
			ws = append(ws, goroutineleak.WorkItem{})
		}
		goroutineleak.ProcessWorkItems(ws)
		time.Sleep(50 * time.Millisecond)
	}
}
