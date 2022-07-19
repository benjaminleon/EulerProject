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
	for _, stripFunc := range []func(string) string{fromLeft, fromRight} {
		truncatedNr := nr
		for {
			truncatedNr = removeDigit(stripFunc, truncatedNr)
			if !primes.IsPrime(truncatedNr) {
				return false
			}
			if truncatedNr < 10 {
				break
			}
		}
	}
	return true
}

func fromLeft(str string) string {
	return str[1:]
}

func fromRight(str string) string {
	return str[:len(str)-1]
}

func removeDigit(stripFunc func(string) string, nr int) int {
	str := fmt.Sprintf("%d", nr)
	truncatedNr, err := strconv.Atoi(stripFunc(str))
	if err != nil {
		panic(err)
	}
	return truncatedNr
}
