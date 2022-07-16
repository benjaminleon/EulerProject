/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in
attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.
We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
There are exactly four non-trivial examples of this type of fraction,
less than one in value, and containing two digits in the numerator and denominator.
If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

package main

import (
	"fmt"
	"strconv"
)

type fraction struct {
	a int
	b int
}

func main() {
	var fractions []fraction
	fr := fraction{a: 1, b: 1}
	for a := 10.0; a < 100; a++ {
		for b := 10.0; b < 100; b++ {
			if isCurious(a, b) {
				fractions = append(fractions, fraction{a: int(a), b: int(b)})
				fr.a *= int(a)
				fr.b *= int(b)
			}
		}
	}
	primes := []int{2, 3, 5, 7, 9, 11, 13, 17, 19} // up until 19 proved to be enough

	for _, prime := range primes {
		for fr.a%prime == 0 && fr.b%prime == 0 {
			fmt.Printf("dividing %v %v with %v\n", fr.a, fr.b, prime)
			fr.a /= prime
			fr.b /= prime
		}
	}
	fmt.Println(fr)
	fmt.Println()
}

func isCurious(a, b float64) bool {
	if a >= b {
		return false
	}
	a1, a2 := split(int(a))
	b1, b2 := split(int(b))
	if a2 == 0 && b2 == 0 {
		return false // trivial
	}
	if a1 == b1 {
		if a2 != 0 && b2 != 0 {
			if a/b == a2/b2 {
				fmt.Printf("Found1 %d/%d\n", int(a), int(b))
				fmt.Printf("a1 a2 b1 b2 %f %f %f %f\n", a1, a2, b1, b2)
				fmt.Printf("%f %f\n", a/b, a2/b2)

			}
		}
	}
	if a1 == b2 {
		if a2 != 0 && b1 != 0 {
			if a/b == a2/b1 {
				fmt.Printf("Found2 %d/%d\n", int(a), int(b))
				return true
			}
		}
	}

	if a2 == b1 {
		if a1 != 0 && b2 != 0 {
			if a/b == a1/b2 {
				fmt.Printf("Found3 %d/%d\n", int(a), int(b))
				return true
			}
		}
	}

	if a2 == b2 {
		if a2 != 0 && b1 != 0 {
			if a/b == a1/b1 {
				fmt.Printf("Found4 %d/%d\n", int(a), int(b))
				return true
			}
		}
	}
	return false
}

func split(nr int) (float64, float64) {
	str := strconv.Itoa(nr)
	leftDigit, _ := strconv.Atoi(string(str[0]))
	rightDigit, _ := strconv.Atoi(string(str[1]))
	return float64(leftDigit), float64(rightDigit)
}
