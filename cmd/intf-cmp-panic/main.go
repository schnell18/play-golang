package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type User struct {
	ID   uint64
	Name string
	Data interface{}
}

func panicFunc1(u User) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic:", r)
		}
		wg.Done()
	}()
	fmt.Println(u == u)
}

func panicFunc2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic:", r)
		}
		wg.Done()
	}()

	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x)
}

func main() {
	user1 := User{
		ID:   66777,
		Name: "Jose",
		Data: "=235252524552454",
	}
	wg.Add(1)
	go panicFunc1(user1)

	user2 := User{
		ID:   66777,
		Name: "Jose",
		Data: []string{"good", "better"},
	}
	wg.Add(1)
	go panicFunc1(user2)
	go panicFunc2()
	wg.Add(1)
	wg.Wait()
}
