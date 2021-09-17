package tempconv

import (
	"testing"
)

func TestCToF(t *testing.T) {
	tests := []struct {
		C Celsius
		F Fahrenheit
	}{
		{-40, -40},
		{0, 32},
		{100, 212},
	}

	for i, test := range tests {
		if actF := CToF(test.C); actF != test.F {
			t.Errorf("Test case #%d failed due to expected: CToF(%s) == %s, actual: %s", i+1, test.C, test.F, actF)
		}
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		F Fahrenheit
		C Celsius
	}{
		{-40, -40},
		{32, 0},
		{212, 100},
	}

	for i, test := range tests {
		if actC := FToC(test.F); actC != test.C {
			t.Errorf("Test case #%d failed due to expected: FToC(%s) == %s, actual: %s", i+1, test.F, test.C, actC)
		}
	}
}

func TestKToC(t *testing.T) {
	tests := []struct {
		K Kelvin
		C Celsius
	}{
		{0, -273.15},
		{1000, -272.15},
	}

	for i, test := range tests {
		if actC := KToC(test.K); actC != test.C {
			t.Errorf("Test case #%d failed due to expected: KToC(%s) == %s, actual: %s", i+1, test.K, test.C, actC)
		}
	}
}

func TestCToK(t *testing.T) {
	tests := []struct {
		C Celsius
		K Kelvin
	}{
		{-273.15, 0},
		{-272.15, 1000},
	}

	for i, test := range tests {
		if actK := CToK(test.C); actK != test.K {
			t.Errorf("Test case #%d failed due to expected: CToK(%s) == %s, actual: %s", i+1, test.C, test.K, actK)
		}
	}
}
