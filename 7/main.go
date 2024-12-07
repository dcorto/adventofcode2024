package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

const day = 7

type Equation struct {
	result    int
	operators []int
}

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

	equations := generateEquations(lines)

	for _, eq := range equations {
		numOperators := len(eq.operators) - 1
		numberOfCombinations := 1 << numOperators

		for i := 0; i < numberOfCombinations; i++ {
			result := eq.operators[0]

			for j := 0; j < numOperators; j++ {
				if result > eq.result {
					break
				}
				pthBit := (1 << j) & i
				if pthBit == 0 {
					result += eq.operators[j+1]
				} else {
					result *= eq.operators[j+1]
				}
			}

			if result == eq.result {
				solution += eq.result
				break
			}
		}
	}

	return solution
}

func generateEquations(lines []string) []Equation {
	var equations []Equation

	// Parse equations
	for _, line := range lines {
		eq := strings.Split(line, ":")
		r, _ := strconv.Atoi(eq[0])
		op, _ := utils.SliceFromStringToInt(strings.Split(strings.Trim(eq[1], " "), " "))
		equations = append(equations, Equation{result: r, operators: op})
	}

	return equations
}

func solutionB() int {
	var solution = 0

	_, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/b.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	return solution
}
