package main

import "fmt"

type Error string

func (e Error) Error() string {
	return string(e)
}

func MakePanic(list []int) int {
	return list[len(list)*2]
}

func ContainPanic2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%#v\n", e)
			// repanic with original panic
			if _, ok := e.(Error); !ok {
				panic(e)
			}
		}
	}()
	list := []int{1, 2, 3, 5}
	MakePanic(list)
}

func ContainPanic1() {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		fmt.Printf("%#v\n", e)
	// 	}
	// }()
	ContainPanic2()
}

func main() {
	ContainPanic1()
}
