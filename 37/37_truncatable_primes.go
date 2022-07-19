/*
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right,
and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/
package main

import (
	"example.com/ben/primes"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sumOfTruncatablePrimes := 0
	cnt := 0
	primeNumbers := primes.GetPrimes(1000000)
	sort.Ints(primeNumbers)

	primeNumbers = primeNumbers[5:] // Remove 2, 3, 5, 7 since they are 1-digit numbers and cannot be truncated

	for _, nr := range primeNumbers {
		if isTruncatablePrime(nr) {
			sumOfTruncatablePrimes += nr
			cnt++
			fmt.Printf("Found %d\n", nr)
		}
	}
	fmt.Printf("The sum of the truncatable primes is %d, and there are %d of them", sumOfTruncatablePrimes, cnt)
}

func isTruncatablePrime(nr int) bool {
	truncatedLeft := nr
	for {
		truncatedLeft = removeDigitFromLeft(truncatedLeft)
		if !primes.IsPrime(truncatedLeft) {
			return false
		}
		if truncatedLeft < 10 {
			break
		}

	}
	truncatedRight := nr
	for {
		truncatedRight = removeDigitFromRight(truncatedRight)
		if !primes.IsPrime(truncatedRight) {
			return false
		}
		if truncatedRight < 10 {
			break
		}
	}
	return true
}

func removeDigitFromLeft(nr int) int {
	str := fmt.Sprintf("%d", nr)
	truncatedNr, err := strconv.Atoi(str[1:])
	if err != nil {
		panic(err)
	}
	return truncatedNr
}

func removeDigitFromRight(nr int) int {
	str := fmt.Sprintf("%d", nr)
	truncatedNr, err := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		panic(err)
	}
	return truncatedNr
}
