package prime

// SievePrime prime number to nth
func SievePrime(nth int64) []int64 {
	primes := make([]int64, nth)
	src := make(chan int64)
	go generate(src)
	var i int64
	for i = 0; i < nth; i++ {
		prime := <-src
		primes[i] = prime
		dst := make(chan int64)
		go filter(src, dst, prime)
		src = dst
	}
	return primes
}

func generate(ch chan<- int64) {
	var i int64
	for i = 2; ; i++ {
		ch <- i
	}
}

func filter(src <-chan int64, dst chan<- int64, prime int64) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}
