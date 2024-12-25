package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

const day = 25

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

	gridOfLocks, gridOfKeys := parseInput(input)

	locks := transform(gridOfLocks)
	keys := transform(gridOfKeys)

	var pairs = 0
	for _, lock := range locks {
		for _, key := range keys {
			if fits(key, lock) {
				pairs++
			}
		}
	}

	solution = pairs
	return solution
}

// fits checks if the key fits the lock
func fits(key []int, lock []int) bool {
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}

// parseInput parses the input into two grids of strings, one for locks and one for keys
func parseInput(input []string) ([][]string, [][]string) {
	var gridOfLocks, gridOfKeys [][]string

	isLock := false
	isKey := false

	var index = 0
	for _, line := range input {
		if line == "" || index == 6 {
			index = 0
			continue
		}

		if line == "#####" && index == 0 {
			isLock = true
			isKey = false
			index++
			continue
		}

		if line == "....." && index == 0 {
			isLock = false
			isKey = true
			index++
			continue
		}

		if isKey {
			gridOfKeys = append(gridOfKeys, strings.Split(line, ""))
		}

		if isLock {
			gridOfLocks = append(gridOfLocks, strings.Split(line, ""))
		}
		index++
	}

	return gridOfLocks, gridOfKeys
}

// transform converts the grid of strings to a grid of integers that represent the height of the keys or locks
func transform(input [][]string) [][]int {
	var output [][]int
	for chunk := range slices.Chunk(input, 5) {
		var temp []int
		var height int
		for i := 0; i < 5; i++ {
			height = 0
			for j := 0; j < 5; j++ {
				if chunk[j][i] == "#" {
					height++
				}
			}
			temp = append(temp, height)
		}
		output = append(output, temp)
	}
	return output
}

func solutionB() string {
	var solution = "\n Ho ho ho!!! \n Thanks again Eric for a fun month of puzzles. \n Merry Christmas and Happy New Year!"
	return solution
}
