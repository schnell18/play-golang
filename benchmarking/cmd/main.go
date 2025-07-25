package main

func main() {
	MemoryWaste(100)
}

func MemoryWaste(n int) []int {
	var data []int
	for i := 0; i < n; i++ {
		data = append(data, i)
	}
	return data
}

func MemoryWaste2(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data = append(data, i)
	}
	return data
}
