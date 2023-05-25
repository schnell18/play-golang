package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	UserID int
	Num    interface {
		int | int8 | int16 | int64 | float32 | float64
	}
)

func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}

func Add[T int | float64](a, b T) T {
	return a + b
}

// tiddle menas type and its alias types
func Add1[T ~int | float64](a, b T) T {
	return a + b
}

func Add2[T constraints.Ordered](a, b T) T {
	return a + b
}

func Add3[T Num](a, b T) T {
	return a + b
}

func main() {
	a, b := 1.01, 9.01
	fmt.Printf("%.2f + %.2f = %.2f \n", a, b, AddFloat(a, b))
	x, y := 1, 9
	u1, u2 := UserID(1000), UserID(2000)
	fmt.Printf("%d + %d = %d \n", x, y, AddInt(x, y))

	fmt.Printf("Generic: %.2f + %.2f = %.2f \n", a, b, Add(a, b))
	fmt.Printf("Generic: %d + %d = %d \n", x, y, Add(x, y))
	fmt.Printf("Generic Type alias: %d + %d = %d \n", u1, u2, Add1(u1, u2))
	fmt.Printf("Generic: %.2f + %.2f = %.2f \n", a, b, Add2(a, b))
	fmt.Printf("Generic: %d + %d = %d \n", x, y, Add2(x, y))
	fmt.Printf("Generic: %.2f + %.2f = %.2f \n", a, b, Add3(a, b))
	fmt.Printf("Generic: %d + %d = %d \n", x, y, Add3(x, y))
}
