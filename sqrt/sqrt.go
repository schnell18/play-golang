package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {

    for _, arg := range os.Args[1:] {
        if f, err := strconv.ParseFloat(arg, 64); err == nil {
            fmt.Println(Sqrt(f))
        } else {
            fmt.Printf("bad input %s\n", arg)
        }
    }
}

func Sqrt(x float64) float64 {
	z0 := 1.2
	z := z0
	for i := 1; ; i += 1 {
		z -= (z*z - x) / (2 * z)
		abs := z - z0
		if z < z0 {
			abs = -abs
		}
		if abs <= 0.00000000000001 {
			// fmt.Printf("%02d = %.15f\n", i, z)
			break
		} else {
			z0 = z
		}
	}
	return z
}
