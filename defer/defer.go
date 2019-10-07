package main

import (
	"fmt"
)

func main() {
    defer fmt.Println("world")
    defer fmt.Println("solar system")
    defer fmt.Println("galaxy")
    defer fmt.Println("universe")
    fmt.Println("Hello")
}
