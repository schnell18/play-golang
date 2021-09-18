package gcd

import (
	"testing"
)

type testData struct {
	X int64
	Y int64
	G int64
}

func TestGCD(t *testing.T) {

	expected := []testData{
		{10, 21, 1},
		{12, 8, 4},
		{20, 15, 5},
		{32, 24, 8},
		{120, 54, 6},
		{121, 54, 1},
		{121, 55, 11},
	}
	for i, exp := range expected {
		if actG := GCD(exp.X, exp.Y); exp.G != actG {
			t.Errorf("%dth test case GCD(%d, %d) == %d != %d", i+1, exp.X, exp.Y, actG, exp.G)
		}
	}
}
