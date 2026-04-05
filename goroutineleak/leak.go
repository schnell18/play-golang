// Package goroutineleak explores the experimental goroutine leak profile in 1.26.
package goroutineleak

type Result struct {
	res WorkResult
	err error
}

type (
	WorkResult struct{}
	WorkItem   struct{}
)

func processWorkItem(_ WorkItem) (WorkResult, error) {
	return WorkResult{}, nil
}

func ProcessWorkItems(ws []WorkItem) ([]WorkResult, error) {
	// Process work items in parallel, aggregating results in ch.
	ch := make(chan Result)
	for _, w := range ws {
		go func() {
			res, err := processWorkItem(w)
			ch <- Result{res, err}
		}()
	}

	// Collect the results from ch, or return an error if one is found.
	var results []WorkResult
	for range len(ws) {
		r := <-ch
		if r.err != nil {
			// This early return may cause goroutine leaks.
			return nil, r.err
		}
		results = append(results, r.res)
	}
	return results, nil
}
