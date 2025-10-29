package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	repeat := 10
	fmt.Println("================UUIDv7============================")
	for i := range repeat {
		id := uuid.Must(uuid.NewV7())
		fmt.Printf("%d =>%s\n", i, id)
	}

	fmt.Println("================UUIDv4============================")
	for i := range repeat {
		id := uuid.New()
		fmt.Printf("%d =>%s\n", i, id)
	}
}
