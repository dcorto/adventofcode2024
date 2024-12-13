package main

import (
	"fmt"
	"os"
	"strings"
)

const day = 13

type Point struct {
	x, y int
}

// Add returns the sum of the vectors p and q.
func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

// Mul returns the multiplication of the vector p by the scalar k.
func (p Point) Mul(k int) Point {
	return Point{p.x * k, p.y * k}
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

	input, _ := os.ReadFile(fmt.Sprintf("%d/input.txt", day))

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		var a, b, c Point
		fmt.Sscanf(s, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&a.x, &a.y, &b.x, &b.y, &c.x, &c.y)
		solution += calc(a, b, c)
	}

	return solution
}

func solutionB() int {
	var solution = 0

	input, _ := os.ReadFile(fmt.Sprintf("%d/input.txt", day))

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		var a, b, c Point
		fmt.Sscanf(s, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&a.x, &a.y, &b.x, &b.y, &c.x, &c.y)
		solution += calc(a, b, c.Add(Point{10000000000000, 10000000000000}))
	}

	return solution
}

// calc returns the number of steps required to reach point c from the origin
func calc(a, b, c Point) int {
	ap := (b.y*c.x - b.x*c.y) / (a.x*b.y - a.y*b.x)
	bp := (a.y*c.x - a.x*c.y) / (a.y*b.x - a.x*b.y)
	if a.Mul(ap).Add(b.Mul(bp)) == c {
		return ap*3 + bp
	}
	return 0
}
