package main

import (
	"fmt"
)

// Solution from 
// https://www.xarg.org/puzzle/project-euler/problem-26/
func main() {

	m := make(map[int][]int)

	for d := 2; d < 1000; d++ {
		decimals := getPattern(d)
		m[d] = decimals
	}

	bestDigit := 0
	for idx, val := range m {
		if len(val) > len(m[bestDigit]) {
			bestDigit = idx
		}
	}
	fmt.Println("Best digit:", bestDigit)
	fmt.Println("Length of repeating pattern:", len(m[bestDigit]))
	fmt.Println("Repeating pattern:", m[bestDigit])

}

func getPattern(denominator int) []int {
	nrOfDecimals := 1000
	nominator := 1
	var remainder int
	m := make(map[int]int)
	decimals := []int{}
	for i := 0; i < nrOfDecimals; i++ {
		remainder = nominator % denominator
		if _, ok := m[remainder]; ok {
			return decimals
		}
		decimals = append(decimals, nominator/denominator)
		m[remainder] = i
		nominator = remainder * 10
	}
	return nil
}

// The solution below found found 97 as the best digit
// func getSmallestCycle(decimals []int, maxNrSlides, maxPatternSize int) []int {
// 	var smallestCycle []int
// 	var left []int
// 	var right []int
// 	for slide := 0; slide <= maxNrSlides; slide++ {
// 		for patternSize := 1; patternSize <= maxPatternSize; patternSize++ {
// 			left = decimals[0:patternSize]
// 			right = decimals[patternSize : patternSize*2]
// 			if testEq(left, right) {
// 				return left
// 			}
// 		}
// 		decimals = decimals[1:]

// 	}
// 	return smallestCycle
// }

// func testEq(a, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i := range a {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
