package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/schnell18/play-golang/prime"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Specify number of prime you want")
		os.Exit(2)
	}
	nth, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad argument %s\n", os.Args[1])
		os.Exit(3)
	}
	primes := prime.SievePrime(nth)
	fmt.Println(primes)
}
