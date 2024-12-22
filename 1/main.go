package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"utils"
)

const day = 1

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() int {
	var solution int = 0

	var left []int
	var right []int

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
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

func solutionB() int {
	var solution int = 0

	var left []int
	var right []int

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
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

	for _, leftItem := range left {
		count := 0
		for _, rightItem := range right {
			if leftItem == rightItem {
				count++
			}
		}

		solution += leftItem * count

	}

	return solution
}
