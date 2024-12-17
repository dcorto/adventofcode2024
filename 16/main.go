package main

import (
	"fmt"
	"slices"
	"utils"
)

const day = 16

// Point represents a point in the map
type Point struct {
	x, y int
}

// Map represents the map
type Map []Row

// Row represents a row in the map
type Row []rune

type Path struct {
	ps    []Point
	score int
}

type Visit struct {
	p Point
	d Point
}
type Travel struct {
	v    Visit
	path Path
}

var WALL = '#'
var START = 'S'
var END = 'E'
var GOOD = 'O'

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

	m := Map{}
	for _, line := range lines {
		m = append(m, Row(line))
	}

	paths := m.Traverse()
	bestScore := bestScore(paths)
	solution = bestScore

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	m := Map{}
	for _, line := range lines {
		m = append(m, Row(line))
	}

	paths := m.Traverse()
	best := bestScore(paths)

	points := map[Point]bool{}
	for _, path := range paths {
		if path.score == best {
			for _, p := range path.ps {
				m.UpdateCell(p, GOOD)
				points[p] = true
			}
		}
	}
	solution = len(points)

	return solution
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

// Draw prints the map (for debugging)
func (m Map) Draw() {
	for _, row := range m {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Print("\n")
	}
}

// getStartAndEndPoint returns the starting end ending points of the map
func (m Map) getStartAndEndPoint() (Point, Point) {
	var s, e Point
	for y, row := range m {
		for x, cell := range row {
			if cell == START {
				s = Point{x, y}
			} else if cell == END {
				e = Point{x, y}
			}
		}
	}
	return s, e
}

// UpdateCell updates the cell at point p with item
func (m Map) UpdateCell(p Point, item rune) {
	m[p.y][p.x] = item
}

// GetCell returns the cell at point p
func (m Map) GetCell(p Point) rune {
	return m[p.y][p.x]
}

func (m Map) Traverse() []Path {
	start, end := m.getStartAndEndPoint()

	queue := []Travel{
		{Visit{start, Point{1, 0}}, Path{[]Point{}, 0}},
		{Visit{start, Point{0, -1}}, Path{[]Point{}, 1000}}, // initial value to avoid going back
	}

	visited := map[Visit]int{}
	var paths []Path
	var current Travel

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		v := current.v
		if visited[v] > 0 && visited[v] < current.path.score {
			continue
		}

		if m.GetCell(v.p) == WALL {
			continue
		}

		newPath := append(slices.Clone(current.path.ps), v.p)
		visited[v] = current.path.score
		if v.p == end {
			paths = append(paths, Path{newPath, current.path.score})
			continue
		}

		cw := Point{v.d.y, -v.d.x}  // clockwise
		ccw := Point{-v.d.y, v.d.x} // counter-clockwise

		queue = append(
			queue,
			Travel{
				Visit{v.p.Add(v.d), v.d}, // straight
				Path{newPath, current.path.score + 1}},
			Travel{
				Visit{v.p.Add(cw), cw}, // clockwise
				Path{newPath, current.path.score + 1001}},
			Travel{
				Visit{v.p.Add(ccw), ccw}, // counter-clockwise
				Path{newPath, current.path.score + 1001}},
		)
	}

	return paths
}

// bestScore returns the best score from the paths
func bestScore(paths []Path) int {
	score := paths[0].score
	for _, path := range paths {
		score = min(score, path.score)
	}
	return score
}
