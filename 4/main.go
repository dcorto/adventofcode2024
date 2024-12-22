package main

import (
	"fmt"
	"strings"
	"utils"
)

const day = 4

var xmas = []string{"X", "M", "A", "S"}

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

var diagonals = [][]int{
	{1, 1},
	{1, -1},
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

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var board [][]string
	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	w := len(board[0])
	h := len(board)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if board[y][x] != xmas[0] {
				continue
			}
			solution += getXMASwordInPosition(x, y, &board, h, w)
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

	board := make([][]byte, 256)
	h := 0

	for _, line := range lines {
		board[h] = make([]byte, len(line))
		copy(board[h], line)
		h++
	}
	w := len(board[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if board[y][x] != byte('A') {
				continue
			}
			if hasXMASInPosition(x, y, &board, h, w) {
				solution += 1
			}
		}
	}

	return solution
}

func getXMASwordInPosition(ox int, oy int, board *[][]string, h int, w int) int {
	count := 0
	for _, d := range directions {
		for i := 0; i <= 3; i++ {
			x := ox + d[0]*i
			y := oy + d[1]*i

			if !isPositionInBoard(x, y, h, w) {
				break
			}

			if (*board)[y][x] != xmas[i] {
				break
			}

			if i == 3 {
				count++
			}
		}
	}
	return count
}

func hasXMASInPosition(ox int, oy int, board *[][]byte, h int, w int) bool {
	r := []int{-1, 1}
	for _, d := range diagonals {
		s := 0
		for _, i := range r {
			x := ox + d[0]*i
			y := oy + d[1]*i
			if !isPositionInBoard(x, y, h, w) {
				return false
			}
			s += int((*board)[y][x])
		}
		if s != int('M')+int('S') {
			return false
		}
	}
	return true
}

func isPositionInBoard(x int, y int, h int, w int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}
