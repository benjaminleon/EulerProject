/*
Take the number 192 and multiply it by each of 1, 2, and 3:
    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)
The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).
What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	largestPan := 0
	for i := 9; i < 10000000; i++ {
		pan := makePan1To9(i)
		if pan > largestPan {
			fmt.Printf("Found a new pandigital number %d from %d\n", pan, i)
			largestPan = pan
		}
	}
	fmt.Println(largestPan)
}

// makePan1To9 returns 0 if it cannot create a 1 to 9 pandigital
func makePan1To9(nr int) int {
	maxDigit := 9
	// Create a 9 digit number
	var attempt string
	i := 1
	for {
		attempt += fmt.Sprintf("%d", nr*i)
		if len(attempt) > maxDigit {
			return 0
		}
		if len(attempt) == maxDigit {
			break
		}

		i++
	}

	// Check if it is pandigital
	for i := 1; i < maxDigit; i++ {
		if !strings.Contains(attempt, fmt.Sprintf("%d", i)) {
			return 0
		}
	}

	// Return it
	pan, err := strconv.Atoi(attempt)
	if err != nil {
		panic(err)
	}
	return pan
}
