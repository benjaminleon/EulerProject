/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once.
For example, 2143 is a 4-digit pandigital and is also prime.
What is the largest n-digit pandigital prime that exists*/

package main

import (
	"example.com/ben/primes"
	"fmt"
	"strings"
)

func main() {
	largest := 0
	for _, p := range primes.GetSortedPrimes(10000000) {
		if isPandigital(p) {
			if p > largest {
				largest = p
			}
		}
	}
	fmt.Printf("The largest pandigital prime is %d", largest)
}

func isPandigital(nr int) bool {
	str := fmt.Sprintf("%d", nr)
	for i := 1; i <= len(str); i++ {
		if !strings.Contains(str, fmt.Sprintf("%d", i)) {
			return false
		}
	}
	return true
}
