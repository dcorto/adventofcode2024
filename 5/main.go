package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

const day = 5

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

	var isUpdate = false

	rules := make([][]int, 0)
	var updates [][]int
	for _, line := range lines {

		if line == "" {
			isUpdate = true
			continue
		}

		if !isUpdate {
			// Parse rules TODO: move to a function
			split := strings.Split(line, "|")
			src, _ := strconv.Atoi(split[0])
			dst, _ := strconv.Atoi(split[1])

			rules = append(rules, []int{src, dst})

			continue
		}

		// Parse updates TODO: move to a function
		var res []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			res = append(res, val)
		}
		updates = append(updates, res)
	}

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			midpoint := getMidpoint(update)
			solution += midpoint
		}
	}

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var isUpdate = false

	rules := make([][]int, 0)
	var updates [][]int
	for _, line := range lines {

		if line == "" {
			isUpdate = true
			continue
		}

		if !isUpdate {
			// Parse rules TODO: move to a function
			split := strings.Split(line, "|")
			src, _ := strconv.Atoi(split[0])
			dst, _ := strconv.Atoi(split[1])

			rules = append(rules, []int{src, dst})

			continue
		}

		// Parse updates TODO: move to a function
		var res []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			res = append(res, val)
		}
		updates = append(updates, res)
	}

	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			// Sort the update using the rules
			sortedLine := topologicalSort(update, rules)
			midpoint := getMidpoint(sortedLine)
			solution += midpoint
		}

	}

	return solution
}

// Function to perform topological sorting for a single update
func topologicalSort(line []int, rules [][]int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for _, rule := range rules {
			// If the current rule is violated, swap the positions of the pages
			for i := 0; i < len(line)-1; i++ {
				for j := i + 1; j < len(line); j++ {
					if line[i] == rule[1] && line[j] == rule[0] {
						// Swap the pages
						line[i], line[j] = line[j], line[i]
						sorted = false
					}
				}
			}
		}
	}
	return line
}

func getMidpoint(update []int) int {
	midpoint := update[len(update)/2]
	return midpoint
}

func isUpdateValid(line []int, rules [][]int) bool {
	for _, rule := range rules {
		aPos, bPos := -1, -1

		// Find positions of both a and b in the line
		for i, page := range line {
			if page == rule[0] {
				aPos = i
			}
			if page == rule[1] {
				bPos = i
			}
		}

		// If both a and b are present and a comes after b, it's invalid
		if aPos != -1 && bPos != -1 && aPos > bPos {
			return false
		}
	}
	return true
}
