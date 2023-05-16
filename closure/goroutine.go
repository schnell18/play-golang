package main

import (
	"fmt"
)

func main() {
	for _, val := range []string{"Math", "Chinese", "English"} {
		go func() {
			fmt.Println(val)
		}()
	}
	// time.Sleep(3 * time.Second)
}
