package main

import "fmt"

func SieveOfEratosthenes(n int) {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	for p := 2; p <= n; p++ {
		if primes[p] {
			fmt.Println(p)
		}
	}
}

func main() {
	var n int = 46
	SieveOfEratosthenes(n)
}
