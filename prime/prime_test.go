package prime

import (
	"testing"
)

func TestFirstTenPrime(t *testing.T) {
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	primes := SievePrime(10)
	if len(primes) != len(expected) {
		t.Errorf(
			"number of generated prime(%d) != expected(%d)",
			len(primes),
			len(expected),
		)
	}
	for i, exp := range expected {
		if primes[i] != exp {
			t.Errorf("%dth prime %d != %d", i, primes[i], exp)
		}
	}
}

func TestFirstPrime(t *testing.T) {
	expected := int64(2)
	primes := SievePrime(1)
	if primes[0] != expected {
		t.Errorf("first prime %d != %d", primes[0], expected)
	}
}

func TestFirstHundredPrimes(t *testing.T) {
	expected := []int64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
		127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
		233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
		283, 293, 307, 311, 313, 317, 331, 337, 347, 349,
		353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463,
		467, 479, 487, 491, 499, 503, 509, 521, 523, 541,
	}
	primes := SievePrime(100)
	if len(primes) != len(expected) {
		t.Errorf(
			"number of generated prime(%d) != expected(%d)",
			len(primes),
			len(expected),
		)
	}
	for i, exp := range expected {
		if primes[i] != exp {
			t.Errorf("%dth prime %d != %d", i, primes[i], exp)
		}
	}
}

func TestPrimesAreAlwaysOddExceptTwo(t *testing.T) {
	primes := SievePrime(50)
	for i, p := range primes {
		if i == 0 {
			if p != 2 {
				t.Errorf("first prime should be 2, got %d", p)
			}
			continue
		}
		if p%2 == 0 {
			t.Errorf("prime at index %d is %d but should be odd", i, p)
		}
	}
}

func TestSievePrimeReturnsEmptyForZero(t *testing.T) {
	primes := SievePrime(0)
	if len(primes) != 0 {
		t.Errorf("expected empty slice for nth=0, got %d primes", len(primes))
	}
}
