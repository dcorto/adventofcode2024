package main

import (
	"fmt"
	"slices"
	"utils"
)

const day = 15

// Point represents a point in the map
type Point struct {
	x, y int
}

// Map represents the map
type Map []Row

// Row represents a row in the map
type Row []rune

var ROBOT = '@'
var BOX = 'O'
var WALL = '#'
var EMPTY = '.'
var LEFTBOX = '['
var RIGHTBOX = ']'

var MOVES = map[rune]Point{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
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

	m, movements := parseInput(lines)

	m.Traverse(movements)
	solution = m.Score()

	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	m, movements := parseInput(lines)

	newMap := m.Resize()
	newMap.Traverse(movements)
	solution = newMap.Score()

	return solution
}

// parseInput returns the map and the movements
func parseInput(lines []string) (Map, []rune) {
	m := Map{}
	var movements []rune
	gridEnd := false

	for _, line := range lines {
		if !gridEnd {
			if len(line) == 0 {
				gridEnd = true
			} else {
				m = append(m, Row(line))
			}
		} else {
			movements = slices.Concat(movements, []rune(line))
		}
	}
	return m, movements
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
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

// getRobotStartPoint returns the starting point of the robot
func (m Map) getRobotStartPoint() Point {
	for y, row := range m {
		for x, cell := range row {
			if ROBOT == cell {
				return Point{x, y}
			}
		}
	}
	panic("Did not find the Robot")
}

// UpdateCell updates the cell at point p with item
func (m Map) UpdateCell(p Point, item rune) {
	m[p.y][p.x] = item
}

// GetCell returns the cell at point p
func (m Map) GetCell(p Point) rune {
	return m[p.y][p.x]
}

// OtherHalf returns the other half of the box
func (m Map) OtherHalf(p Point) Point {
	if m.GetCell(p) == LEFTBOX {
		return p.Add(MOVES['>'])
	} else {
		return p.Add(MOVES['<'])
	}
}

func (m Map) CanMove(p1, p2 Point) bool {
	switch m.GetCell(p2) {
	case EMPTY:
		return true
	case WALL:
		return false
	default: // LEFTBOX, RIGHTBOX
		d := p2.Sub(p1)
		if d == MOVES['<'] || d == MOVES['>'] {
			return m.CanMove(p2, p2.Add(d))
		} else {
			otherHalf := m.OtherHalf(p2)
			return m.CanMove(p2, p2.Add(d)) && m.CanMove(otherHalf, otherHalf.Add(d))
		}
	}
}

func (m Map) Move(p1, p2 Point) Point {
	cellContent := m.GetCell(p1)
	switch m.GetCell(p2) {
	case EMPTY:
		m.UpdateCell(p2, cellContent)
		m.UpdateCell(p1, EMPTY)
		return p2
	case WALL:
		return p1
	case BOX:
		if p2 != m.Move(p2, p2.Add(p2.Sub(p1))) {
			m.UpdateCell(p2, cellContent)
			m.UpdateCell(p1, EMPTY)
			return p2
		} else {
			return p1
		}
	default: // LEFTBOX, RIGHTBOX
		d := p2.Sub(p1)
		if d == MOVES['<'] || d == MOVES['>'] {
			if p2 != m.Move(p2, p2.Add(d)) {
				m.UpdateCell(p2, cellContent)
				m.UpdateCell(p1, EMPTY)
				return p2
			} else {
				return p1
			}
		} else {
			otherHalf := m.OtherHalf(p2)
			if m.CanMove(p2, p2.Add(d)) && m.CanMove(otherHalf, otherHalf.Add(d)) {
				m.Move(p2, p2.Add(d))
				m.Move(otherHalf, otherHalf.Add(d))
				m.UpdateCell(p1, EMPTY)
				m.UpdateCell(p2, cellContent)
				m.UpdateCell(otherHalf, EMPTY)
				return p2
			} else {
				return p1
			}
		}
	}
}

func (m Map) Traverse(movements []rune) {
	p := m.getRobotStartPoint()
	//m.Draw()
	for _, movement := range movements {
		p = m.Move(p, p.Add(MOVES[movement]))
		//m.Draw()
	}
}

func (m Map) Resize() Map {
	newMap := make(Map, len(m))
	for y, row := range m {
		newMap[y] = make(Row, len(row)*2)
		for x, cell := range row {
			switch cell {
			case EMPTY:
				newMap[y][x*2] = EMPTY
				newMap[y][x*2+1] = EMPTY
			case BOX:
				newMap[y][x*2] = LEFTBOX
				newMap[y][x*2+1] = RIGHTBOX
			case WALL:
				newMap[y][x*2] = WALL
				newMap[y][x*2+1] = WALL
			case ROBOT:
				newMap[y][x*2] = ROBOT
				newMap[y][x*2+1] = EMPTY
			}
		}
	}
	return newMap
}

func (m Map) Score() int {
	score := 0
	for y, row := range m {
		for x, cell := range row {
			if cell == BOX || cell == LEFTBOX {
				score += x + y*100
			}
		}
	}
	return score
}
