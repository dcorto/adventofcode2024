package main

import (
	"fmt"
	"strings"
	"utils"
)

const day = 19

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	towels, designs := parseInput(lines)
	solution = getPossibleCountA(towels, designs)

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	towels, designs := parseInput(lines)
	solution = getPossibleCountB(towels, designs)

	return solution
}

// isTowelPossible returns the number of possible designs of towels
func isTowelPossible(towel string, towels []string, possible map[string]int) int {
	if val, ok := possible[towel]; ok {
		return val
	}

	isPossibleCount := 0

	for _, t := range towels {
		if len(t) > len(towel) {
			continue
		}
		if strings.Index(towel, t) == 0 {
			if len(t) == len(towel) {
				isPossibleCount++
				continue
			}
			isPossibleCount += isTowelPossible(towel[len(t):], towels, possible)
		}
	}
	possible[towel] = isPossibleCount
	return isPossibleCount
}

// getPossibleCount returns the number of possible designs of towels
func getPossibleCountA(towels []string, designs []string) int {
	var count = 0
	possible := make(map[string]int)
	for _, d := range designs {
		isPossibleCount := isTowelPossible(d, towels, possible)
		if isPossibleCount > 0 {
			count += 1
		}
	}
	return count
}

// getPossibleCount returns the number of designs that can be made with the towels
func getPossibleCountB(towels []string, designs []string) int {
	var count = 0
	possible := make(map[string]int)
	for _, d := range designs {
		isPossibleCount := isTowelPossible(d, towels, possible)
		count += isPossibleCount
	}
	return count
}

// parseInput returns two slices, one with the towels and one with the designs
func parseInput(input []string) (towels []string, designs []string) {
	for _, row := range input {

		if len(row) == 0 {
			continue
		}

		if strings.Contains(row, ",") {
			towels = append(towels, strings.Split(row, ",")...)
		} else {
			designs = append(designs, strings.TrimSpace(row))
		}
	}

	for i, t := range towels {
		towels[i] = strings.TrimSpace(t)
	}

	return towels, designs
}
