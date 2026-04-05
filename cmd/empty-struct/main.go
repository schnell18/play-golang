package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var emptyInt1 int
	var emptyInt2 int
	var emptyString1 string
	var emptyString2 string
	var emptyStruct1 struct{}
	var emptyAny any

	fmt.Printf("Size of empty int: %d\n", unsafe.Sizeof(emptyInt1))
	fmt.Printf("Size of empty string: %d\n", unsafe.Sizeof(emptyString1))
	fmt.Printf("Size of empty struct: %d\n", unsafe.Sizeof(emptyStruct1))
	fmt.Printf("Size of empty any: %d\n", unsafe.Sizeof(emptyAny))

	var emptyStruct2 struct{}
	fmt.Printf("mem address of empty int1: %p\n", &emptyInt1)
	fmt.Printf("mem address of empty int2: %p\n", &emptyInt2)
	fmt.Printf("mem address of empty string1: %p\n", &emptyString1)
	fmt.Printf("mem address of empty string2: %p\n", &emptyString2)
	fmt.Printf("mem address of empty struct 1: %p\n", &emptyStruct1)
	fmt.Printf("mem address of empty struct 2: %p\n", &emptyStruct2)
	fmt.Printf("mem address of empty any: %p\n", &emptyAny)
}
