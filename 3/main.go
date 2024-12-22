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
	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
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
	var isActive = true
	lines, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))

	for _, line := range lines {
		result, err := extractOnlyEnabledAndValidMultiplicationValueFromLine(line, &isActive)
		if err != nil {
			fmt.Println("Error:", err)
			return solution
		}
		solution += result
	}

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

func extractOnlyEnabledAndValidMultiplicationValueFromLine(line string, isActive *bool) (int, error) {
	var result = 0
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := regex.FindAllStringSubmatch(line, -1)

	for _, match := range matches {

		if match[0] == "do()" {
			*isActive = true
			continue
		}

		if match[0] == "don't()" {
			*isActive = false
			continue
		}

		if *isActive {
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
	}

	return result, nil
}
