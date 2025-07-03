package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 3}

	x = append(x, 8)
	// shows how to append slice using explode operator
	x = append(x, y...)

	fmt.Printf("%v\n", x)
	fmt.Printf("%q\n", x)
	fmt.Printf("%+v\n", x)
	fmt.Printf("%#v\n", x)
}
