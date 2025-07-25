package main

import (
	"fmt"
	"testing"
)

func BenchmarkMemoryWaste(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoryWaste(100)
	}
}

func BenchmarkMemoryWaste2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoryWaste2(100)
	}
}

func BenchmarkMemoryWasteBLoop(b *testing.B) {
	for b.Loop() {
		MemoryWaste(100)
	}
}

func BenchmarkMemoryWasteBLoop2(b *testing.B) {
	for b.Loop() {
		MemoryWaste2(100)
	}
}

func BenchmarkDeadCodeEliminationBLoop(b *testing.B) {
	fmt.Println("B.Loop()")
	for b.Loop() {
		isCond(201)
	}
}

func BenchmarkDeadCodeElimination(b *testing.B) {
	fmt.Println("B.N")
	for i := 0; i < b.N; i++ {
		isCond(201)
	}
}

func BenchmarkDeadCodeElimination2(b *testing.B) {
	ret := false
	for i := 0; i < b.N; i++ {
		ret = isCond(201)
	}
	_ = ret
}

func isCond(b byte) bool {
	if b%3 == 1 && b%7 == 2 && b%17 == 11 && b%31 == 9 {
		return true
	}
	return true
}
