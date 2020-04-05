package functions

import "fmt"

func Max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("max() requires at least one argument")
	}
	maxVal := vals[0]
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}

func Min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("max() requires at least one argument")
	}
	minVal := vals[0]
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal, nil
}
