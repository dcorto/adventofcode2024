package main

import (
	"fmt"
	"strings"
	"utils"
)

const day = 9

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	//solutionB := solutionB()
	//fmt.Println("Solution B:", solutionB)
}

func solutionA() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		fmt.Println("Error:", err)
		return solution
	}

	var diskMap []int
	for _, line := range lines {
		diskMap, _ = utils.SliceFromStringToInt(strings.Split(line, ""))
	}

	//generate list of diskBlocks from diskMap
	var diskBlocks [][]int
	var counter = 0
	var currentBlock []int

	for p, block := range diskMap {
		var freeSpace = false

		if (p % 2) == 1 {
			freeSpace = true
		}

		var i = 0
		for i < block {
			if freeSpace {
				currentBlock = append(currentBlock, -1) // -1 means free space
			} else {
				currentBlock = append(currentBlock, counter)
			}
			i++
		}

		if !freeSpace {
			counter++
		}
	}

	diskBlocks = append(diskBlocks, currentBlock)

	//move diskBlocks to the right (defragmentation)
	for _, blocks := range diskBlocks {
		var j = 0
		for i := len(blocks) - 1; i >= 0; i-- { // from right to left
			file := blocks[i]

			if j >= i {
				break
			}

			if file != -1 {
				for j = 0; j < len(blocks); j++ {

					if j >= i {
						break
					}

					if blocks[j] == -1 { //move file to new position
						blocks[j] = file
						blocks[i] = -1
						break
					}
				}
			}

		}
	}

	//calculate checksum
	for _, blocks := range diskBlocks {
		var counter = 0
		for i, block := range blocks {
			if block != -1 {
				counter += i * block
			}
		}
		solution += counter
	}

	return solution
}

func solutionB() int {
	var solution = 0
	return solution
}
