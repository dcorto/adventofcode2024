package main

import (
	"fmt"
	"math"
	"utils"
)

const day = 20

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

	lines, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))

	grid, start := parseInput(lines)

	dist := bfs(start, grid)

	for p1 := range dist {
		for p2 := range dist {
			d := int(math.Abs(float64(p2.x)-float64(p1.x)) + math.Abs(float64(p2.y)-float64(p1.y)))
			if d <= 20 && dist[p2] >= dist[p1]+d+100 {
				if d <= 2 {
					solution++
				}
			}
		}
	}
	return solution
}

func parseInput(lines []string) (map[Point]rune, Point) {
	grid, start := map[Point]rune{}, Point{}
	for y, s := range lines {
		for x, r := range s {
			if r == 'S' {
				start = Point{x, y}
			}
			grid[Point{x, y}] = r
		}
	}
	return grid, start
}

func solutionB() int {
	var solution = 0

	lines, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))

	grid, start := parseInput(lines)

	dist := bfs(start, grid)

	for p1 := range dist {
		for p2 := range dist {
			d := int(math.Abs(float64(p2.x)-float64(p1.x)) + math.Abs(float64(p2.y)-float64(p1.y)))
			if d <= 20 && dist[p2] >= dist[p1]+d+100 {
				solution++
			}
		}
	}

	return solution
}

func bfs(start Point, grid map[Point]rune) map[Point]int {
	queue, dist := []Point{start}, map[Point]int{start: 0}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := p.Add(d)
			if _, ok := dist[n]; !ok && grid[n] != '#' {
				queue, dist[n] = append(queue, n), dist[p]+1
			}
		}
	}
	return dist
}
