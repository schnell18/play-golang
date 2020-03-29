package fibonacci

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	cases := []struct {
		in, want uint64
	}{
		{0, 0}, {1, 1},
		{2, 1}, {3, 2},
		{4, 3}, {5, 5},
		{6, 8}, {7, 13},
		{8, 21}, {9, 34},
		{10, 55}, {11, 89},
		{12, 144}, {13, 233},
		{14, 377}, {15, 610},
	}
	for _, c := range cases {
		got := Fibonacci(c.in)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestFirstTenFibonacci(t *testing.T) {
	expected := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, exp := range expected {
		act := Fibonacci(uint64(i))
		if exp != act {
			t.Error(fmt.Sprintf("Calcualted %dth fibonacci number %d != %d", i, act, exp))
		}
	}
}

func TestSecondTenFibonacci(t *testing.T) {
	expected := []uint64{55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	for i, exp := range expected {
		act := Fibonacci(uint64(i + 10))
		if exp != act {
			t.Error(fmt.Sprintf("Calcualted %dth fibonacci number %d != %d", i, act, exp))
		}
	}
}
