package main

import "fmt"

type CustomMap[K comparable, V int | string] map[K]V

func main() {
	m := make(CustomMap[int, string])
	m[3] = "3"
	m2 := make(CustomMap[string, string])
	m2["name"] = "Justin"
	m2["title"] = "Mr"
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", m2)
}
