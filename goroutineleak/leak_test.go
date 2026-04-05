package goroutineleak

import (
	"log"
	"testing"
)

func BenchmarkProcessWorkItems(b *testing.B) {
	for b.Loop() {
		items := []WorkItem{
			{}, {}, {}, {}, {}, {}, {}, {},
		}
		if _, err := ProcessWorkItems(items); err != nil {
			log.Fatalf("fail to process work items due to %v", err)
		}
	}
}
