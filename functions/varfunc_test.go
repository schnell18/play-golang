package functions

import (
	"fmt"
	"testing"
)

func TestMaxInsufficientArgs(t *testing.T) {
	_, err := Max()
	if err == nil {
		t.Error(fmt.Sprintf("Max() should return error"))
	}
}

func TestMaxGood(t *testing.T) {
	exp := 55
	vals := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	act, _ := Max(vals...)
	if exp != act {
		t.Error(fmt.Sprintf("Max(%v) returns %d shoud be %d", vals, act, exp))
	}
}

func TestMinInsufficientArgs(t *testing.T) {
	_, err := Min()
	if err == nil {
		t.Error(fmt.Sprintf("Min() should return error"))
	}
}

func TestMinGood(t *testing.T) {
	exp := 0
	vals := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	act, _ := Min(vals...)
	if exp != act {
		t.Error(fmt.Sprintf("Min(%v) returns %d shoud be %d", vals, act, exp))
	}
}
