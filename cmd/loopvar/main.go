package main

import (
	"fmt"
	"time"
)

// As of 2023-07-23, to get the intended result, run this program as follow
// gvm install go1.21rc3
// gvm use go1.21rc3
// export GOEXPERIMENT=loopvar
// go run main.go
// gvm can be installed from https://github.com/moovweb/gvm
func main() {
	facts := []string{"alpha", "beta", "gamma"}

	for _, fact := range facts {
		go func() {
			fmt.Println(fact)
		}()
	}

	time.Sleep(3 * time.Second)
}
