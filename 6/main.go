package main

import (
	"fmt"
	"strings"
	"utils"
)

const day = 6

const guard = "^"

type Position struct {
	x int
	y int
}

var directions = [4][2]int{
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
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

	board := make([][]string, 0)

	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	path := getPath(board)
	solution = len(path)
	return solution
}

func getPath(board [][]string) []Position {
	width := len(board[0])
	height := len(board)
	visited := make([]bool, width*height)
	var path []Position

	position := startPosition(board)
	dir := 0 //up

	for {
		idx := position.x*width + position.y
		if !visited[idx] {
			visited[idx] = true
			path = append(path, position)
		}

		nextPos := Position{
			x: position.x + directions[dir][0],
			y: position.y + directions[dir][1],
		}

		if isGuardOutsideBoard(nextPos, width, height) {
			return path
		}

		if board[nextPos.y][nextPos.x] == "#" {
			dir = (dir + 1) & 3
		} else {
			position = nextPos
		}
	}
}

func isGuardOutsideBoard(nextPos Position, width int, height int) bool {
	return nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height
}

func startPosition(board [][]string) Position {
	for y, row := range board {
		for x, cell := range row {
			if cell == guard {
				return Position{x, y}
			}
		}
	}
	return Position{}
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	board := make([][]string, 0)

	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}

	solution = countLoopPositions(board, getPath(board))
	return solution
}

func countLoopPositions(board [][]string, initialPath []Position) int {
	start := startPosition(board)
	count := 0
	for _, pos := range initialPath {
		if board[pos.y][pos.x] == "." {
			board[pos.y][pos.x] = "#"
			if hasLoop(board, start) {
				count++
			}
			board[pos.y][pos.x] = "."
		}
	}
	return count
}

func hasLoop(board [][]string, start Position) bool {
	width := len(board[0])
	height := len(board)
	visited := make([]uint8, width*height)
	pos := start
	dir := 0 // up

	for {
		idx := pos.y*width + pos.x
		dirBit := uint8(1 << dir)

		if visited[idx]&dirBit != 0 {
			return true
		}
		visited[idx] |= dirBit

		nextPos := Position{
			x: pos.x + directions[dir][0],
			y: pos.y + directions[dir][1],
		}

		if isGuardOutsideBoard(nextPos, width, height) {
			return false
		}

		if board[nextPos.y][nextPos.x] == "#" {
			dir = (dir + 1) & 3
		} else {
			pos = nextPos
		}
	}
}

// printBoard only for debug purposes
func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}
