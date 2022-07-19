/*
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package main

import "fmt"

func main() {
	sumOfPalindromes := 0
	for i := 1; i < 1000000; i++ {
		b := fmt.Sprintf("%b", i)
		d := fmt.Sprintf("%d", i)
		if isPalindrome(d) && isPalindrome(b) {
			sumOfPalindromes += i
		}
	}
	fmt.Printf("The sum of the double based palindromes is %d", sumOfPalindromes)
}

func isPalindrome(number string) bool {
	for i := 0; i < len(number)/2; i++ {
		if string(number[i]) != string(number[len(number)-1-i]) {
			return false
		}
	}
	return true
}
