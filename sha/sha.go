package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Printf("sha256: %x\n", sha256.Sum256([]byte("hello world")))
	fmt.Printf("sha1: %x\n", sha1.Sum([]byte("hello world")))
}
