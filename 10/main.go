package main

import (
	"fmt"
	"strings"
	"utils"
)

const day = 10

type point struct {
	x int
	y int
}

func (p *point) add(other *point) point {
	return point{p.x + other.x, p.y + other.y}
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

	var grid [][]int
	for _, line := range lines {
		l, _ := utils.SliceFromStringToInt(strings.Split(line, ""))
		grid = append(grid, l)
	}

	for _, trailhead := range findTrailheads(&grid) {
		solution += findTrails(&grid, trailhead, false)
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

	var grid [][]int
	for _, line := range lines {
		l, _ := utils.SliceFromStringToInt(strings.Split(line, ""))
		grid = append(grid, l)
	}

	for _, trailhead := range findTrailheads(&grid) {
		solution += findTrails(&grid, trailhead, true)
	}
	return solution
}

// findTrails finds all the possible trails given a trailhead
func findTrails(grid *[][]int, trailhead point, uniquePaths bool) int {
	trails := 0
	visited := map[point]bool{}
	paths := []point{trailhead}
	var currentPoint point

	width, height := len((*grid)[0])-1, len(*grid)-1

	for len(paths) > 0 {
		currentPoint, paths = paths[0], paths[1:]

		if !uniquePaths && visited[currentPoint] {
			continue
		}

		visited[currentPoint] = true

		currentHeight := (*grid)[currentPoint.y][currentPoint.x]
		if currentHeight == 9 {
			trails++
			continue
		}

		var dirs = []point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		for _, d := range dirs {
			neighborPoint := currentPoint.add(&d)
			if isPointInGrid(&neighborPoint, width, height) && (*grid)[neighborPoint.y][neighborPoint.x] == currentHeight+1 {
				paths = append(paths, neighborPoint)
			}
		}
	}
	return trails
}

// findTrailheads finds all the points where a trail starts (0)
func findTrailheads(grid *[][]int) []point {
	var trailheads []point
	for y, row := range *grid {
		for x, cell := range row {
			if cell == 0 {
				trailheads = append(trailheads, point{x, y})
			}
		}
	}
	return trailheads
}

// isPointInGrid checks if a point is inside the grid
func isPointInGrid(p *point, width, height int) bool {
	return p.x >= 0 && p.y >= 0 && p.x <= width && p.y <= height
}
