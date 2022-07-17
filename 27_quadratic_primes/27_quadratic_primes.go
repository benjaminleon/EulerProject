package main

import (
	"example.com/ben/primes"
	"fmt"
)

func main() {
	best := best{}

	for a := -1000; a <= 1000; a++ {
		for _, b := range primes.GetPrimes(1000) {
			n := 0
			for {
				quad := n*n + a*n + b
				if !primes.IsPrime(quad) {
					break
				}
				//fmt.Printf("a: %v, b: %v, n: %v, quad: %v\n", a, b, n, quad)

				if n > best.n {
					best.a = a
					best.b = b
					best.n = n
					best.quad = quad
				}
				n++
			}
		}
	}

	fmt.Printf("Best a: %v, b: %v, n: %v, quad: %v\n", best.a, best.b, best.n, best.quad)
}

type best struct {
	a, b, n, quad int
}
