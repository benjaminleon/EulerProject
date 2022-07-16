package main

import (
	"fmt"
	"strconv"
)

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
Find the sum of all numbers which are equal to the sum of the factorial of their digits.
Note: As 1! = 1 and 2! = 2 are not sums they are not included.
*/

func main() {
	sum := 0
	for i := 10; i < 1000000; i++ {
		digits := split(i)
		factorialDigitSum := 0
		for _, digit := range digits {
			factorialDigitSum += factorial(digit)
		}
		if factorialDigitSum == i {
			sum += i
			fmt.Printf("Found %d\n", i)
		}
	}
	fmt.Printf("The answer is %d", sum)
}

func split(nr int) []int {
	var digits []int
	str := strconv.Itoa(nr)
	for i := 0; i < len(str); i++ {
		digit, _ := strconv.Atoi(string(str[i]))
		digits = append(digits, digit)
	}
	return digits
}

func factorial(nr int) (result int) {
	result = 1
	for i := 2; i <= nr; i++ {
		result *= i
	}
	return
}
