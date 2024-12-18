package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

const day = 18

var XMAX, YMAX = 70, 70
var BYTES = 1024
var DIRECTIONS = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Visit struct {
	pos  Point
	path []Point
}

// Point represents a point in the map
type Point struct {
	x, y int
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
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

	positions := parseInput(lines)
	path := findPath(positions[0:BYTES])
	solution = len(path) - 1
	return solution
}

func solutionB() string {
	var solution = "0,0"

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	positions := parseInput(lines)
	brokenPathByte := searchFirstBrokenPath(positions)
	solution = strconv.Itoa(brokenPathByte.x) + "," + strconv.Itoa(brokenPathByte.y)

	return solution
}

// findPath returns the path from the start to the end
func findPath(fallingBytes []Point) []Point {
	start, end := Point{0, 0}, Point{XMAX, YMAX}
	positions := []Visit{{start, []Point{}}}
	visited := map[Point]bool{}
	fallen := map[Point]bool{}

	for _, f := range fallingBytes {
		fallen[f] = true
	}

	var visit Visit

	for len(positions) > 0 {
		visit, positions = positions[0], positions[1:]

		if visit.pos == end {
			return append(visit.path, visit.pos)
		}
		if visited[visit.pos] {
			continue
		}
		visited[visit.pos] = true

		for _, d := range DIRECTIONS {
			np := visit.pos.Add(d)
			if inBounds(np) && !fallen[np] {
				positions = append(positions, Visit{np, append(slices.Clone(visit.path), visit.pos)})
			}
		}
	}
	return []Point{}
}

// searchFirstBrokenPath returns the first point where the path is broken
func searchFirstBrokenPath(positions []Point) Point {
	min, max := BYTES, len(positions)-1
	for min <= max {
		center := (min + max) / 2
		path := findPath(positions[0 : center+1])
		if len(path) > 0 {
			min = center + 1
		} else {
			max = center - 1
		}
	}
	return positions[min]
}

// inBounds returns true if the point is within the bounds of the memory space (XMAX, YMAX)
func inBounds(pos Point) bool {
	return pos.x >= 0 && pos.x <= XMAX && pos.y >= 0 && pos.y <= YMAX
}

// parseInput returns a slice of Points from the input
func parseInput(lines []string) []Point {
	var positions []Point
	for _, line := range lines {
		pair := strings.Split(line, ",")
		positions = append(positions, Point{stringToInt(pair[0]), stringToInt(pair[1])})
	}
	return positions
}

// stringToInt converts a string to an integer
func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
