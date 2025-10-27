package main

import "fmt"

func main() {
	indexOutOfBound()
	nilDereference()
	typeAssertion()
	fmt.Println("all panic handled, exit normally")
}

func indexOutOfBound() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()
	// case 1: index out-of-bound
	nums := []int{}
	fmt.Println(nums[1])
}

func nilDereference() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()
	// case 2: dereference nil pointer
	var str *string
	fmt.Println(*str)
}

func typeAssertion() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()
	// case 3: type assertion
	var str any = "123"
	_ = str.(int)
}

// func writeUninitialisedMap() {
// 	// case 3: write uninitialised map
// 	var h map[string]int
// 	h[""] = 1
// }
