package main

import (
	"fmt"
	"regexp"
	"strconv"
	"utils"
)

const day = 3

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() int {
	var solution = 0
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/a.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		result, err := extractValidMultiplicationValueFromLine(line)
		if err != nil {
			fmt.Println("Error:", err)
			return solution
		}
		solution += result
	}

	return solution
}

func solutionB() int {
	var solution = 0
	return solution
}

func extractValidMultiplicationValueFromLine(line string) (int, error) {
	var result = 0
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		op1, err := strconv.Atoi(match[1])
		if err != nil {
			return result, err
		}

		op2, err := strconv.Atoi(match[2])
		if err != nil {
			return result, err
		}

		result += op1 * op2
	}

	return result, nil
}
