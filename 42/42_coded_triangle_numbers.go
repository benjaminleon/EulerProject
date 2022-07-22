/*
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.
Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/ben/ProjectEuler/42/words.txt")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	words, _ := reader.ReadAll()
	count := 0
	triangleNumbers := getTriangleNumbers(300) // Ten Z's would make a word value of 260

	for _, word := range words[0] {
		wordValue := 0
		for _, letter := range word {
			wordValue += int(letter) - 64
		}
		if _, ok := triangleNumbers[wordValue]; ok {
			count++
			fmt.Printf("Found triangle word %s\n", word)
		}
	}

	fmt.Println("Number of triangle words", count)
}

func getTriangleNumbers(last int) map[int]bool {
	var numbers = map[int]bool{}
	for n := 1; n < last; n++ {
		numbers[n*(n+1)/2] = true
	}
	return numbers
}
