package main

import "testing"

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
		got := Fib(c.in)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
