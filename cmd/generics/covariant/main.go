package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range values {
		newValues = append(newValues, mapFunc(v))
	}
	return newValues
}

func Map[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValues = append(newValues, mapFunc(v))
	}
	return newValues
}

func main() {
	ints := []int{1, 2, 3, 5}
	fmt.Printf("%v\n", MapValues(ints, func(i int) int { return i * 2 }))
	fmt.Printf("generics: %v\n", Map(ints, func(i int) int { return i * 2 }))
}
