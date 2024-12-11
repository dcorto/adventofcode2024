package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

const day = 11

type cacheKey struct {
	stone int
	n     int
}

var cache = map[cacheKey]int{}

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

	var stones []int
	stones, _ = utils.SliceFromStringToInt(strings.Split(lines[0], " "))

	for _, stone := range stones {
		solution += blink(stone, 25)
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

	var stones []int
	stones, _ = utils.SliceFromStringToInt(strings.Split(lines[0], " "))

	for _, stone := range stones {
		solution += blink(stone, 75)
	}

	return solution
}

func blink(stone int, n int) int {
	if n == 0 {
		return 1
	}

	key := cacheKey{stone, n}
	val, found := cache[key]
	if found {
		return val
	}

	var result int
	str := strconv.Itoa(stone)
	if stone == 0 {
		result = blink(1, n-1)
	} else if len(str)%2 == 0 {
		result = blink(atoi(str[0:(len(str)/2)]), n-1) + blink(atoi(str[len(str)/2:]), n-1)
	} else {
		result = blink(stone*2024, n-1)
	}

	cache[key] = result
	return result
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
