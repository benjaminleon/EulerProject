package main

import (
	"fmt"
	"math"
)

func main() {
	// Best a: -41, b: 461, n: 102, quad: 6683 was wrong

	best := best{}

	for a := -1000; a <= 1000; a++ {
		for _, b := range getPrimes(1000) {
			n := 0
			for {
				quad := n*n + a*n + b
				if !isPrime(quad) {
					break
				}
				fmt.Printf("a: %v, b: %v, n: %v, quad: %v\n", a, b, n, quad)
				// fmt.Println("is prime")

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

func getPrimes(n int) []int {
	if n < 2 {
		return []int{1}
	}
	primes := map[int]bool{}

	numbers := make([]int, n) // [1, 2, 3, 4, 5, 6, 7, 8, 9...]
	for i := 1; i < n; i++ {
		numbers[i-1] = i
	}
	// fmt.Println(primes)

	for _, number := range numbers[1:] {
		val, ok := primes[number]
		if val == false && ok == true { // This has already been marked as a non-prime
			continue
		}
		primes[number] = true
		multiple := 2
		for {
			if multiple*number >= n {
				break
			} else {
				primes[multiple*number] = false
			}
			multiple++
		}
	}

	// fmt.Println(primes)
	ps := []int{}
	for key, val := range primes {
		if val == true {
			ps = append(ps, key)
		}
	}

	// sort.Ints(ps)
	return ps
}

func isPrime(n int) bool {
	if n < 1 {
		return false
	}
	// fmt.Printf("Checking if %v is a prime\n", n)
	sqrtN := int(math.Floor(math.Sqrt(float64(n))))
	// fmt.Println("sqrtN:", sqrtN)
	for _, p := range getPrimes(sqrtN) {
		if n%p == 0 {
			return false
		}
	}
	return true
}
