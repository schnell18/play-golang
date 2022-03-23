package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Printf("Sizeof(x)   = %d  Alignof(x)   = %d\n", unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Printf("Sizeof(x.a) = %d  Alignof(x.a) = %d  Offsetof(x.a) = %d\n", unsafe.Sizeof(x.a), unsafe.Alignof(x), unsafe.Offsetof(x.a))
	fmt.Printf("Sizeof(x.b) = %d  Alignof(x.b) = %d  Offsetof(x.b) = %d\n", unsafe.Sizeof(x.b), unsafe.Alignof(x), unsafe.Offsetof(x.b))
	fmt.Printf("Sizeof(x.c) = %d  Alignof(x.c) = %d  Offsetof(x.c) = %d\n", unsafe.Sizeof(x.c), unsafe.Alignof(x), unsafe.Offsetof(x.c))
}
