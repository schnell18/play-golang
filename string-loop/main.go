package main

import "fmt"

type MyString string

func (m MyString) String() string {
	// return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
	return fmt.Sprintf("MyString=%v", m) // Error: will recur forever.
	// return fmt.Sprintf("MyString=%s", string(m))
}

func init() {
	fmt.Println("Init 1")
}

func init() {
	fmt.Println("Init 2")
}

func main() {
	// var s1 MyString = "forever loop"
	s1 := "good"
	fmt.Println(s1)
}

func init() {
	fmt.Println("Init 3")
}
