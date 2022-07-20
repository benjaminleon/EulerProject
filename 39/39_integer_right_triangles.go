/*
If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
{20,48,52}, {24,45,51}, {30,40,50}
For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

import (
	"fmt"
	"math"
)

// Warning, this takes ~20 seconds to complete
func main() {
	solutions := map[int]int{}
	for perimeter := 120; perimeter < 1000; perimeter++ {
		solutions[perimeter] = 0
		for a := 1; a < perimeter; a++ {
			for b := 1; b < a; b++ {
				for c := 1; c <= perimeter-a-b; c++ {
					if a+b+c == perimeter && isRightAngle(a, b, c) {
						solutions[perimeter]++
						//fmt.Printf("Found a solution with perimeter %d with {%d, %d, %d}\n", perimeter, a, b, c)
					}
				}
			}
		}
	}
	bestPerimeter := 0
	mostSolutions := 0
	for perimeter, nrSolutions := range solutions {
		if nrSolutions > mostSolutions {
			mostSolutions = nrSolutions
			bestPerimeter = perimeter
		}
	}
	fmt.Printf("The best perimeter is %d, with %d solutions\n", bestPerimeter, mostSolutions)
}

func isRightAngle(a, b, c int) bool {
	return math.Abs(math.Sqrt(math.Pow(float64(a), 2)+math.Pow(float64(b), 2))-float64(c)) < 0.000001
}
