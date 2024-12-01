package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"utils"
)

func main() {
	fmt.Println("Solution for Day 1")
	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)
}

func solutionA() int {
	var solution int = 0

	var left []int
	var right []int

	lines, err := utils.ReadLinesFromFile("1/a.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	for _, line := range lines {
		slice := strings.Split(line, "   ")

		t, err := strconv.Atoi(slice[0])
		if err == nil {
			left = append(left, t)
		}

		t, err = strconv.Atoi(slice[1])
		if err == nil {
			right = append(right, t)
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	for i, leftItem := range left {
		rightItem := right[i]
		dist := int(math.Abs(float64(leftItem) - float64(rightItem)))
		solution += dist
	}

	return solution
}
