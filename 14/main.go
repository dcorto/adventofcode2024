package main

import (
	"fmt"
	"strconv"
	"utils"
)

const day = 14

const wide = 101
const tall = 103

type Robot struct {
	px, py, vx, vy int
}

func (r Robot) Move() Robot {
	var x, y int

	x = r.px + r.vx
	y = r.py + r.vy

	if x >= wide || x < 0 {
		if r.vx > 0 {
			x = x - wide
		} else {
			x = x + wide
		}
	}

	if y >= tall || y < 0 {
		if r.vy > 0 {
			y = y - tall
		} else {
			y = y + tall
		}
	}

	return Robot{x, y, r.vx, r.vy}
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

	var robots []Robot

	for _, line := range lines {
		var r Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)
		robots = append(robots, r)
	}

	for i := 0; i < 100; i++ {
		for j, r := range robots {
			robots[j] = r.Move()
		}
	}

	grid := positionRobots(robots)

	q1 := q1(grid)
	g2 := q2(grid)
	q3 := q3(grid)
	q4 := q4(grid)

	solution = q1 * g2 * q3 * q4

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var robots []Robot

	for _, line := range lines {
		var r Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)
		robots = append(robots, r)
	}

	var seconds = 0
	for {

		for j, r := range robots {
			robots[j] = r.Move()
		}
		seconds++

		if checkOverlap(robots) == false {
			break
		}
	}

	solution = seconds

	return solution
}

func positionRobots(robots []Robot) [tall][wide]int {
	var grid [tall][wide]int

	for _, r := range robots {
		x := r.px
		y := r.py

		if &grid[y][x] == nil {
			grid[y][x] = 1
		} else {
			n := grid[y][x]
			n++
			grid[y][x] = n
		}
	}

	return grid
}

func q1(grid [tall][wide]int) int {
	var sum int

	for y := 0; y < tall/2; y++ {
		for x := 0; x < wide/2; x++ {
			sum += grid[y][x]
		}
	}

	return sum
}

func q2(grid [tall][wide]int) int {
	var sum int

	for y := 0; y < tall/2; y++ {
		for x := wide / 2; x < wide; x++ {
			sum += grid[y][x]
		}
	}

	return sum
}

func q3(grid [tall][wide]int) int {
	var sum int

	for y := (tall / 2) + 1; y < tall; y++ {
		for x := 0; x < wide/2; x++ {
			sum += grid[y][x]
		}
	}

	return sum
}

func q4(grid [tall][wide]int) int {
	var sum int

	for y := (tall / 2) + 1; y < tall; y++ {
		for x := (wide / 2) + 1; x < wide; x++ {
			sum += grid[y][x]
		}
	}

	return sum
}

func checkOverlap(robots []Robot) bool {
	var grid [tall][wide]string

	for x := 0; x < wide; x++ {
		for y := 0; y < tall; y++ {
			grid[y][x] = "."
		}
	}

	for _, r := range robots {
		x := r.px
		y := r.py

		if grid[y][x] == "." {
			grid[y][x] = "1"
		} else {
			return true
		}
	}

	return false
}

// printRobots prints the robots in a grid (only for debug)
func printRobots(robots []Robot) {
	var grid [tall][wide]string

	for x := 0; x < wide; x++ {
		for y := 0; y < tall; y++ {
			grid[y][x] = "."
		}
	}

	for _, r := range robots {
		x := r.px
		y := r.py

		if grid[y][x] == "." {
			grid[y][x] = "1"
		} else {
			s := grid[y][x]
			n, _ := strconv.Atoi(s)
			n++
			grid[y][x] = strconv.Itoa(n)
		}
	}

	for _, gridRow := range grid {
		fmt.Println(gridRow)
	}
}
