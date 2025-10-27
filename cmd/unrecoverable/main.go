package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(5)
	switch n {
	case 0:
		concurrentMapReadWrite()
	case 1:
		outOfMemory()
	case 2:
		stackOverflow()
	case 3:
		nilFunc()
	default:
		deadLock()
	}

	fmt.Println("all panic handled, exit normally")
}

func concurrentMapReadWrite() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()

	h := map[string]int{}

	go func() {
		for {
			h["x"] = 1
		}
	}()

	for {
		_ = h["y"]
	}
}

func outOfMemory() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()
	_ = make([]int64, 1<<40)
}

func stackOverflow() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()
	var fun func([1000]int64)
	fun = func(i [1000]int64) {
		fun(i)
	}
	fun([1000]int64{})
}

func nilFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()
	var fun func()
	go fun()
}

func deadLock() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught err: ", err)
		}
	}()

	<-make(chan int)
}
