package main

import "testing"

func TestSqrt(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{2.0 , 1.414213562373095},
		 {4.0, 2.0},
		 {1.0, 1.0},
		 {100.0, 10.0},
		 {400.0, 20.0},
	}
	for _, c := range cases {
		got := Sqrt(c.in)
		if got != c.want {
			t.Errorf("Sqrt(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}
