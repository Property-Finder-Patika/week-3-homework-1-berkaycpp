package main

import (
	"fmt"
	"math"
)

func PrimeFactors(n int) {
	for n%2 == 0 {
		fmt.Println(2)
		n = n / 2
	}

	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			fmt.Println(i)
			n = n / i
		}
	}
	if n > 2 {
		fmt.Println(n)
	}
}

func main() {
	var n int = 615
	PrimeFactors(n)
}
