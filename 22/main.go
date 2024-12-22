package main

import (
	"fmt"
	"utils"
)

const day = 22

// Sequence represents a 4 consecutive differences
type Sequence [4]int

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() int {
	var solution = 0

	input, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	numbers, _ := utils.SliceFromStringToInt(input)

	for _, value := range numbers {
		for i := 0; i < 2000; i++ {
			value = generateSecret(value)
		}
		solution += value
	}

	return solution
}

func solutionB() int {
	var solution = 0

	input, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	numbers, _ := utils.SliceFromStringToInt(input)

	// Map to store pattern totals
	total := make(map[Sequence]int)

	// Process each initial value from a number in the input
	for _, x := range numbers {
		last := x % 10
		sequenceList := make([][2]int, 0, 2000)

		// Generate pattern list through evolution
		for i := 0; i < 2000; i++ {
			x = generateSecret(x)
			temp := x % 10
			sequenceList = append(sequenceList, [2]int{temp - last, temp})
			last = temp
		}

		// Track seen sequences using a map for O(1) lookup
		seen := make(map[Sequence]bool)

		// Check sequence patterns of length 4 throughout the list
		for i := 0; i < len(sequenceList)-4; i++ {
			// Build sequence pattern from consecutive differences
			var sequence Sequence
			for j := 0; j < 4; j++ {
				sequence[j] = sequenceList[i+j][0]
			}

			val := sequenceList[i+3][1]

			// If pattern hasn't been seen in this iteration
			if !seen[sequence] {
				seen[sequence] = true
				total[sequence] += val
			}
		}
	}

	// Find maximum value in totals
	maxVal := 0
	for _, v := range total {
		if v > maxVal {
			maxVal = v
		}
	}
	solution = maxVal

	return solution
}

// generateSecret calculates the next secret number in the sequence using multiplication, division, mixing (XOR),
// and pruning (modulo) operations
func generateSecret(secret int) int {
	// Step 1: Multiply by 64, mix, and prune
	result := secret * 64
	secret = mixValue(secret, result)
	secret = pruneValue(secret)

	// Step 2: Divide by 32, mix, and prune
	result = secret / 32
	secret = mixValue(secret, result)
	secret = pruneValue(secret)

	// Step 3: Multiply by 2048, mix, and prune
	result = secret * 2048
	secret = mixValue(secret, result)
	secret = pruneValue(secret)

	return secret
}

// mixValue performs a bitwise XOR operation between the secret and the given value
func mixValue(secret, value int) int {
	return secret ^ value
}

// pruneValue calculates the value of the secret modulo 16777216
func pruneValue(secret int) int {
	return secret % 16777216
}
