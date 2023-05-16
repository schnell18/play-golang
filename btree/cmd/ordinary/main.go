package main

import "github.com/tidwall/btree"

const (
	size = 1_00_000_000
)

func main() {
	var tr btree.Set[int]

	for i := 0; i < size; i++ {
		tr.Insert(i)
	}
}
