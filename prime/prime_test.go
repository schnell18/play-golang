package prime

import (
	"fmt"
	"testing"
)

func TestFirstTenPrime(t *testing.T) {
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	primes := SievePrime(10)
	if len(primes) != len(expected) {
		t.Error(fmt.Sprintf("number of generated prime(%d) != expected(%d)", len(primes), len(expected)))
	}
	for i, exp := range expected {
		if primes[i] != exp {
			t.Error(fmt.Sprintf("%dth prime %d != %d", i, primes[i], exp))
		}
	}
}
