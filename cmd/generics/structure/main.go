package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type User struct {
	ID   int
	Name string
	Data interface{}
}

type User2[T CustomData] struct {
	ID   int
	Name string
	Data T
}

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

func main() {
	u1 := User2[int32]{
		ID:   13434,
		Name: "Justin",
		Data: 11,
	}
	fmt.Printf("%v\n", u1)
	u2 := User2[float32]{
		ID:   13434,
		Name: "Justin",
		Data: 11.1,
	}
	fmt.Printf("%v\n", u2)
}
