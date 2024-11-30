package main

import (
	"fmt"
	"utils"
)

func main() {
	fmt.Println("Solution for Day 1")
	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)
}

func solutionA() int {
	var solution int = -1
	lines, err := utils.ReadLinesFromFile("1/a.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return solution
}
