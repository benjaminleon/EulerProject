package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0

	a, b, c, d, e, f := 0, 0, 0, 0, 0, 2
	for {
		if f+10*e+100*d+1000*c+10000*b+100000*a == power(a)+power(b)+power(c)+power(d)+power(e)+power(f) {
			fmt.Printf("Found %d%d%d%d%d%d\n", a, b, c, d, e, f)
			sum += f + 10*e + 100*d + 1000*c + 10000*b + 100000*a
		}

		f++
		e += f / 10
		d += e / 10
		c += d / 10
		b += c / 10
		a += b / 10

		f %= 10
		e %= 10
		d %= 10
		c %= 10
		b %= 10
		a %= 10

		if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 && f == 0 {
			break
		}

	}

	fmt.Printf("Sum %d", sum)
}

func power(num int) int {
	return int(math.Pow(float64(num), 5))
}
