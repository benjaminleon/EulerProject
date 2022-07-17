/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
How many circular primes are there below one million?
*/
package main

import (
	"example.com/ben/primes"
	"fmt"
	"strconv"
)

func main() {
	nrCircularPrimes := 0
	for _, p := range primes.GetPrimes(1000000) {
		r := getRotations(p)
		success := true
		for _, n := range r {
			if !primes.IsPrime(n) {
				success = false
				break
			}
		}
		if success {
			nrCircularPrimes++
		}
	}
	fmt.Printf("There are %d circular primes", nrCircularPrimes)
}

// getRotations takes e.g. 123 and returns [231, 312]
func getRotations(nr int) []int {
	var result []int
	str := strconv.Itoa(nr)
	for i := 1; i < len(str); i++ {
		rotatedNr, err := strconv.Atoi(str[i:] + str[:i])
		if err != nil {
			panic(err)
		}
		result = append(result, rotatedNr)
	}

	return result
}
