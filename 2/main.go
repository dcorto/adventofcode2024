package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"utils"
)

const day = 2

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

	var distance int
	for _, line := range lines {
		var isSafe = true
		var isIncreasing = false

		slice := strings.Split(line, " ")

		first, err := strconv.Atoi(slice[0])
		if err != nil {
			fmt.Println("Error:", err)
			return solution
		}

		second, err := strconv.Atoi(slice[1])
		if err != nil {
			fmt.Println("Error:", err)
			return solution
		}

		if (first - second) > 0 {
			isIncreasing = true
		}

		for i := 1; i < len(slice); i++ {
			current, err := strconv.Atoi(slice[i])
			if err != nil {
				fmt.Println("Error:", err)
				return solution
			}

			before, err := strconv.Atoi(slice[i-1])
			if err != nil {
				fmt.Println("Error:", err)
				return solution
			}

			if before-current < 0 && isIncreasing {
				isSafe = false
				//fmt.Println("Not safe because increasing")
				break
			}

			if before-current > 0 && !isIncreasing {
				isSafe = false
				//fmt.Println("Not safe because decreasing")
				break
			}

			distance = int(math.Abs(float64(current) - float64(before)))
			if distance == 0 || distance > 3 {
				isSafe = false
				//fmt.Println("Not safe because distance")
				break
			}
		}

		if isSafe {
			solution++
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

	for _, line := range lines {
		report, err := utils.SliceFromStringToInt(strings.Split(line, " "))
		if err != nil {
			fmt.Println("Error:", err)
			return solution
		}

		isIncreasing := isIncreasingReport(report)
		isSafe := isSafeReport(report, isIncreasing)
		if isSafe {
			solution++
		}

		if !isSafe {
			l := len(report)
			for i := 0; i < l; i++ {
				newReport := slices.Clone(report)
				newReport = append(newReport[:i], newReport[i+1:]...)
				isIncreasingN := isIncreasingReport(newReport)
				isSafeN := isSafeReport(newReport, isIncreasingN)
				if isSafeN {
					solution++
					break
				}
			}
		}
	}

	return solution
}

func isIncreasingReport(report []int) bool {
	if (report[0] - report[1]) > 0 {
		return true
	}
	return false
}

func isSafeReport(report []int, isIncreasing bool) bool {
	var isSafe = true

	//fmt.Println(report)

	for i := 1; i < len(report); i++ {
		current := report[i]
		before := report[i-1]

		if before-current < 0 && isIncreasing {
			isSafe = false
			//fmt.Println("Not safe because increasing")
			break
		}

		if before-current > 0 && !isIncreasing {
			isSafe = false
			//fmt.Println("Not safe because decreasing")
			break
		}

		distance := int(math.Abs(float64(current) - float64(before)))
		if distance == 0 || distance > 3 {
			isSafe = false
			//fmt.Println("Not safe because distance")
			break
		}
	}

	return isSafe
}
