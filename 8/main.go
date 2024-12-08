package main

import (
	"fmt"
	"strings"
	"utils"
)

const day = 8

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

	board := make([][]string, 0)
	mapAntenas := make(map[string][][]int)

	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	for y, row := range board {
		for x, cell := range row {
			if cell == "." {
				continue
			}
			mapAntenas[cell] = append(mapAntenas[cell], []int{x, y})
		}
	}

	width := len(board[0])
	height := len(board)

	allSignals := make([][]int, 0)
	for _, points := range mapAntenas {
		if len(points) > 1 {
			for i := 0; i < len(points); i++ {
				for j := i + 1; j < len(points); j++ {
					signalPoints := addAntinodesForPosition(points[i], points[j], width, height)
					for _, signal := range signalPoints {
						allSignals = appendIfUnique(allSignals, signal)
					}
				}
			}
		}
	}

	solution = len(allSignals)

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	board := make([][]string, 0)
	mapAntenas := make(map[string][][]int)

	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	for y, row := range board {
		for x, cell := range row {
			if cell == "." {
				continue
			}
			mapAntenas[cell] = append(mapAntenas[cell], []int{x, y})
		}
	}

	width := len(board[0])
	height := len(board)

	allSignals := make([][]int, 0)
	for _, points := range mapAntenas {
		if len(points) > 1 {
			for i := 0; i < len(points); i++ {
				for j := i + 1; j < len(points); j++ {
					signalPoints := addExtendedAntinodesForPosition(points[i], points[j], width, height)
					for _, signal := range signalPoints {
						allSignals = appendIfUnique(allSignals, signal)
					}
				}
			}
		}
	}

	solution = len(allSignals)

	return solution
}

func addAntinodesForPosition(src1 []int, src2 []int, width int, height int) [][]int {
	var result [][]int

	point1 := make([]int, 2)
	point2 := make([]int, 2)

	deltaX := src2[0] - src1[0]
	deltaY := src2[1] - src1[1]

	point1[0] = src1[0] - deltaX
	point1[1] = src1[1] - deltaY

	point2[0] = src2[0] + deltaX
	point2[1] = src2[1] + deltaY

	if width > point1[0] && point1[0] >= 0 && width > point1[1] && point1[1] >= 0 {
		result = append(result, point1)
	}

	if height > point2[0] && point2[0] >= 0 && height > point2[1] && point2[1] >= 0 {
		result = append(result, point2)
	}

	return result
}

func addExtendedAntinodesForPosition(src1 []int, src2 []int, width int, height int) [][]int {
	var result [][]int

	deltaX := src2[0] - src1[0]
	deltaY := src2[1] - src1[1]

	x, y := src1[0], src1[1]
	for {
		x -= deltaX
		y -= deltaY
		if x < 0 || x >= width || y < 0 || y >= height {
			break
		}
		result = append(result, []int{x, y})
	}

	x, y = src2[0], src2[1]
	for {
		x += deltaX
		y += deltaY
		if x < 0 || x >= width || y < 0 || y >= height {
			break
		}
		result = append(result, []int{x, y})
	}

	result = append(result, src1)
	result = append(result, src2)

	return result
}

func appendIfUnique(signals [][]int, signal []int) [][]int {
	for _, s := range signals {
		if s[0] == signal[0] && s[1] == signal[1] {
			return signals
		}
	}

	return append(signals, signal)
}
