package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

const day = 12

type region struct {
	area      int
	perimeter int
	positions map[string]struct{}
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

	var garden [][]string
	for _, line := range lines {
		garden = append(garden, strings.Split(line, ""))
	}

	regions := calculateRegions(garden)
	for _, r := range regions {
		solution += r.area * r.perimeter
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

	var garden [][]string
	for _, line := range lines {
		garden = append(garden, strings.Split(line, ""))
	}

	regions := calculateRegions(garden)
	for _, r := range regions {
		solution += r.area * getSides(r.positions)
	}

	return solution
}

func calculateRegions(garden [][]string) []region {
	var regions []region
	checked := make(map[string]struct{})
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			if _, ok := checked[getPositionKey(x, y)]; !ok {
				r := region{
					positions: make(map[string]struct{}),
				}
				checked[getPositionKey(x, y)] = struct{}{}
				getRegion(garden, x, y, checked, &r)
				regions = append(regions, r)
			}

		}
	}
	return regions
}

func getRegion(garden [][]string, x, y int, checked map[string]struct{}, region *region) {
	region.area++
	region.perimeter += 4
	region.positions[getPositionKey(x, y)] = struct{}{}
	if x > 0 && garden[y][x-1] == garden[y][x] {
		if _, ok := checked[getPositionKey(x-1, y)]; !ok {
			checked[getPositionKey(x-1, y)] = struct{}{}
			getRegion(garden, x-1, y, checked, region)
		}
		region.perimeter--
	}
	if x < len(garden[0])-1 && garden[y][x+1] == garden[y][x] {
		if _, ok := checked[getPositionKey(x+1, y)]; !ok {
			checked[getPositionKey(x+1, y)] = struct{}{}
			getRegion(garden, x+1, y, checked, region)
		}
		region.perimeter--
	}
	if y > 0 && garden[y-1][x] == garden[y][x] {
		if _, ok := checked[getPositionKey(x, y-1)]; !ok {
			checked[getPositionKey(x, y-1)] = struct{}{}
			getRegion(garden, x, y-1, checked, region)
		}
		region.perimeter--
	}
	if y < len(garden)-1 && garden[y+1][x] == garden[y][x] {
		if _, ok := checked[getPositionKey(x, y+1)]; !ok {
			checked[getPositionKey(x, y+1)] = struct{}{}
			getRegion(garden, x, y+1, checked, region)
		}
		region.perimeter--
	}
}

func getPositionKey(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func getSides(region map[string]struct{}) int {
	sides := 0
	for i := range region {
		position := strings.Split(i, ":")
		x, _ := strconv.Atoi(position[0])
		y, _ := strconv.Atoi(position[1])
		_, left := region[getPositionKey(x-1, y)]
		_, right := region[getPositionKey(x+1, y)]
		_, up := region[getPositionKey(x, y-1)]
		_, down := region[getPositionKey(x, y+1)]
		_, upperLeft := region[getPositionKey(x-1, y-1)]
		_, upperRight := region[getPositionKey(x+1, y-1)]
		_, downLeft := region[getPositionKey(x-1, y+1)]
		_, downRight := region[getPositionKey(x+1, y+1)]

		if !left && !up {
			sides++
		}
		if !right && !up {
			sides++
		}
		if !left && !down {
			sides++
		}
		if !right && !down {
			sides++
		}
		if !upperRight && up && right {
			sides++
		}
		if !upperLeft && up && left {
			sides++
		}
		if !downLeft && down && left {
			sides++
		}
		if !downRight && down && right {
			sides++
		}
	}
	return sides
}
