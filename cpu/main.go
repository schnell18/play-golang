package main

import (
	"fmt"
	"runtime"
)

func init() {
	// runtime.GOMAXPROCS(int(0.8 * float64(runtime.NumCPU())))
	runtime.GOMAXPROCS(int(0.7 * float64(runtime.NumCPU())))
}

func main() {
	fmt.Printf("Physical CPU cores: %d\n", runtime.NumCPU())
	fmt.Printf("User Specified CPU cores: %d\n", runtime.GOMAXPROCS(0))
}
