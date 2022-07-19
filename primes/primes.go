package primes

import "math"

func IsPrime(n int) bool {
	if n < 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	// fmt.Printf("Checking if %v is a prime\n", n)
	sqrtN := int(math.Floor(math.Sqrt(float64(n))))
	// fmt.Println("sqrtN:", sqrtN)
	for _, p := range GetPrimes(sqrtN) { // [2, 3, 5, 7, 11...]
		if n%p == 0 {
			return false
		}
	}
	return true
}

// GetPrimes generates a list of prime numbers that are less than or equal to n
func GetPrimes(n int) []int {
	if n < 2 {
		return []int{1}
	}
	primes := map[int]bool{}

	numbers := make([]int, n) // [1, 2, 3, 4, 5, 6, 7, 8, 9...]
	for i := 1; i <= n; i++ {
		numbers[i-1] = i
	}
	// fmt.Println(primes)

	for _, number := range numbers[1:] {
		// fmt.Printf("Looking at %v: \n", number)
		val, ok := primes[number]
		if val == false && ok == true { // This has already been marked as a non-prime
			continue
		}
		primes[number] = true
		multiple := 2
		for {
			if multiple*number >= n {
				break
			} else {
				primes[multiple*number] = false
			}
			multiple++
		}
	}

	// fmt.Println(primes)
	ps := []int{}
	for key, val := range primes {
		if val == true {
			ps = append(ps, key)
		}
	}

	// sort.Ints(ps)
	return ps
}
