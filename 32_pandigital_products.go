/*

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	products := map[int]bool{}
	for multiplicand := 1; multiplicand < 100; multiplicand++ {
		for multiplier := 1; multiplier < 3000; multiplier++ {
			if isPandigital(multiplicand, multiplier) {
				if !products[multiplicand*multiplier] { // Don't print duplicates
					fmt.Printf("%d * %d = %d\n", multiplicand, multiplier, multiplicand*multiplier)
				}
				products[multiplicand*multiplier] = true
			}
		}
	}
	sum := 0
	for prod, _ := range products {
		sum += prod
	}
	fmt.Printf("Sum of the products: %d\n", sum)
}

func isPandigital(nr1, nr2 int) bool {
	m := map[rune]bool{}
	str := strconv.Itoa(nr1) + strconv.Itoa(nr2) + strconv.Itoa(nr1*nr2)
	if len(str) != 9 || strings.Contains(str, "0") {
		return false
	}

	// See if any digit occurs more than once
	for _, digit := range str {
		m[digit] = true
	}
	return len(m) == len(str)
}
