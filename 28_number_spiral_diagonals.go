package main

import "fmt"

func main() {

	field := generateEmptyField(1001)
	field = fillField(field)

	fmt.Printf("Sum of diagonal: %d\n", diagonalSum(field))
}

func generateEmptyField(size int) [][]int {
	field := make([][]int, size)
	for i := 0; i < size; i++ {
		field[i] = make([]int, size)
	}
	return field
}

func printField(field [][]int) {
	for i := range field {
		for j := range field[i] {
			fmt.Printf("%v  ", field[i][j])
		}
		fmt.Println()
	}
}

func fillField(field [][]int) [][]int {
	size := len(field)
	middle := size / 2
	fmt.Printf("size: %v. middle: %v\n", size, middle)
	finalNr := size * size
	stepLength := 1
	// First take 1 step right, then 1 down, then 2 left and 2 up, then 3 right and 3 down
	firstTimeWithStepSize := true
	direction := "right"
	field[middle][middle] = 1
	i, j := middle, middle
	nr := 2
	fmt.Printf("Final number: %v\n", finalNr)
	done := false
	for {
		step := 0
		for step < stepLength {
			i, j = takeStep(i, j, direction)
			field[i][j] = nr
			step++
			nr++
			if nr > finalNr {
				done = true
				break
			}
		}
		if done {
			break
		}
		direction = changeDirection(direction)
		if !firstTimeWithStepSize {
			stepLength++
		}
		firstTimeWithStepSize = !firstTimeWithStepSize

	}
	return field
}

func diagonalSum(field [][]int) int {
	size := len(field)
	sum := 0

	// Sum diagonal from top left to bottom right
	i, j := 0, 0
	for {
		sum += field[i][j]
		i++
		j++
		if i >= size {
			break
		}
	}

	// Sum diagonal from botton left to top right
	i = size - 1
	j = 0
	for {
		sum += field[i][j]
		i--
		j++
		if i < 0 {
			break
		}
	}

	// Remove the middle which has been counted twice
	return sum - 1
}

func changeDirection(d string) string {
	switch d {
	case "right":
		{
			return "down"
		}
	case "down":
		{
			return "left"
		}
	case "left":
		{
			return "up"
		}
	case "up":
		{
			return "right"
		}
	}
	panic("Invalid direction")
}

func takeStep(i int, j int, direction string) (int, int) {
	switch direction {
	case "right":
		{
			return i, j + 1
		}
	case "down":
		{
			return i + 1, j
		}
	case "left":
		{
			return i, j - 1
		}
	case "up":
		{
			return i - 1, j
		}
	}
	panic("wrong direction")
}
