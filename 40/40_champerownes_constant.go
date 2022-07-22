/*
An irrational decimal fraction is created by concatenating the positive integers:
0.123456789101112131415161718192021...
It can be seen that the 12th digit of the fractional part is 1.
If dn represents the nth digit of the fractional part, find the value of the following expression.
d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	answer := 1
	count := 0
	nextIdx := 1
	for i := 1; i < 1000000; i++ {
		// Optimize by checking larger (more likely) range first
		// This proved however to be fast enough :)
		if i < 10 {
			count += 1
		} else if i < 100 {
			count += 2
		} else if i < 1000 {
			count += 3
		} else if i < 10000 {
			count += 4
		} else if i < 100000 {
			count += 5
		} else if i < 1000000 {
			count += 6
		}

		if count >= nextIdx {
			overshoot := count - nextIdx
			str := strconv.Itoa(i)
			digit, err := strconv.Atoi(string(str[len(str)-1-overshoot]))
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d overshot idx %d with %d, taking %d\n", i, nextIdx, overshoot, digit)
			answer *= digit
			nextIdx *= 10
		}

	}
	fmt.Printf("The answer is %d\n", answer)
}
