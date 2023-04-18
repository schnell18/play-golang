package fib

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fibonacci Sequence Test Suite")
}

var _ = Describe("Fibonacci sequence", func() {

	Context("Small integers", func() {
		It("0th Fibonacci number", func() {
			Expect(Fibonacci(0)).Should(BeZero())
		})
		It("1th Fibonacci number", func() {
			Expect(Fibonacci(1)).Should(BeEquivalentTo(1))
		})
		It("2th Fibonacci number", func() {
			Expect(Fibonacci(2)).Should(BeEquivalentTo(1))
		})
		It("3th Fibonacci number", func() {
			Expect(Fibonacci(3)).Should(BeEquivalentTo(2))
		})
	})

	Context("Big integers", func() {
		It("12th Fibonacci number", func() {
			Expect(Fibonacci(12)).Should(Equal(uint64(144)))
		})
		It("13th Fibonacci number", func() {
			Expect(Fibonacci(13)).Should(Equal(uint64(233)))
		})
		It("15th Fibonacci number", func() {
			Expect(Fibonacci(15)).Should(Equal(uint64(610)))
		})
	})

	Context("Continuous integers", func() {
		cases := []struct {
			arg  uint64
			want uint64
		}{
			{0, 0}, {-1, 1}, {2, 1}, {3, 2},
			{4, 3}, {5, 5}, {6, 8}, {7, 13},
			{8, 21}, {9, 34}, {10, 55}, {11, 89},
			{12, 144}, {13, 233}, {14, 377}, {15, 610},
			{16, 987}, {17, 1597},
		}
		for _, tc := range cases {
			It(fmt.Sprintf("%dth Fibonacci number", tc.arg), func() {
				Expect(Fibonacci(tc.arg)).Should(Equal(tc.want))
			})
		}
	})

})
