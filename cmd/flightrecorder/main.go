package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
	"time"
)

type bucket struct {
	mu      sync.Mutex
	guesses int
}

func main() {
	fr := trace.NewFlightRecorder(trace.FlightRecorderConfig{
		MinAge:   200 * time.Millisecond,
		MaxBytes: 1 << 20, // 1 MiB
	})
	fr.Start()

	buckets := make([]bucket, 100)

	// 每分钟发送报告
	go func() {
		for range time.Tick(10 * time.Second) {
			sendReport(buckets)
		}
	}()

	answer := rand.Intn(len(buckets))

	http.HandleFunc(
		"/guess-number",
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			guess, err := strconv.Atoi(r.URL.Query().Get("guess"))
			if err != nil || !(0 <= guess && guess < len(buckets)) {
				http.Error(w, "invalid 'guess' value", http.StatusBadRequest)
				return
			}

			b := &buckets[guess]
			b.mu.Lock()
			b.guesses++
			b.mu.Unlock()

			_, _ = fmt.Fprintf(
				w,
				"guess: %d, correct: %t",
				guess,
				guess == answer,
			)

			log.Printf(
				"HTTP request: endpoint=/guess-number guess=%d duration=%s",
				guess,
				time.Since(start),
			)
			if fr.Enabled() && time.Since(start) > 6*time.Microsecond {
				go captureSnapshot(fr)
			}
		},
	)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func sendReport(buckets []bucket) {
	counts := make([]int, len(buckets))

	for index := range buckets {
		b := &buckets[index]
		b.mu.Lock()
		defer b.mu.Unlock()

		counts[index] = b.guesses
	}

	b, err := json.Marshal(counts)
	if err != nil {
		log.Printf("failed to marshal report data: error=%s", err)
		return
	}

	url := "http://localhost:8091/guess-number-report"
	if _, err := http.Post(url, "application/json", bytes.NewReader(b)); err != nil {
		log.Printf("failed to send report: %s", err)
	}
}

var once sync.Once

func captureSnapshot(fr *trace.FlightRecorder) {
	once.Do(func() {
		f, err := os.Create("snapshot.trace")
		if err != nil {
			log.Printf("opening snapshot file %s failed: %s", f.Name(), err)
			return
		}
		defer f.Close()

		_, err = fr.WriteTo(f)
		if err != nil {
			log.Printf("writing snapshot to file %s failed: %s", f.Name(), err)
			return
		}

		fr.Stop()
		log.Printf("captured a flight recorder snapshot to %s", f.Name())
	})
}
