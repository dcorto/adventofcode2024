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

	var diskMap []int
	for _, line := range lines {
		diskMap, _ = utils.SliceFromStringToInt(strings.Split(line, ""))
	}

	//generate list of diskBlocks from diskMap
	diskBlocks := generateDiskBlocksFromDiskMap(diskMap)

	//move files to the right (defragmentation)
	var j = 0
	for i := len(diskBlocks) - 1; i >= 0; i-- { // from right to left
		file := diskBlocks[i]

		if j >= i {
			break
		}

		if file != -1 {
			for j = 0; j < len(diskBlocks); j++ {

				if j >= i {
					break
				}

				if diskBlocks[j] == -1 { //move file to new position
					diskBlocks[j] = file
					diskBlocks[i] = -1
					break
				}
			}
		}
	}

	//calculate checksum
	solution = calculateChecksum(diskBlocks)

	return solution
}

func solutionB() int {
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
	diskBlocks := generateDiskBlocksFromDiskMap(diskMap)

	//move files to the right (defragmentation by ID file number)
	currentFile := -1
	currentFileLength := 0
	for i := len(diskBlocks) - 1; i > 0; i-- {
		if diskBlocks[i] != -1 {
			currentFileLength += 1
			currentFile = i
			if diskBlocks[i] != diskBlocks[i-1] {
				if currentFileLength != 0 {
					moveFile(diskBlocks, currentFileLength, currentFile)
					currentFileLength = 0
				}
			}
		} else {
			if currentFileLength != 0 {
				moveFile(diskBlocks, currentFileLength, currentFile)
				currentFileLength = 0
			}
		}
	}

	//calculate checksum
	solution = calculateChecksum(diskBlocks)

	return solution
}

func generateDiskBlocksFromDiskMap(diskMap []int) []int {
	var diskBlocks []int
	var counter = 0

	for p, block := range diskMap {
		var freeSpace = false

		if (p % 2) == 1 {
			freeSpace = true
		}

		var i = 0
		for i < block {
			if freeSpace {
				diskBlocks = append(diskBlocks, -1) // -1 means free space
			} else {
				diskBlocks = append(diskBlocks, counter)
			}
			i++
		}

		if !freeSpace {
			counter++
		}
	}
	return diskBlocks
}

func moveFile(fileBlock []int, length int, originalStartIndex int) {
	freeSpaceCount := 0
	freeSpaceStartIndex := -1
	for i := 1; i <= originalStartIndex; i++ {
		if fileBlock[i-1] != -1 && fileBlock[i] == -1 {
			if freeSpaceCount >= length {
				writeFile(fileBlock, fileBlock[originalStartIndex], length, freeSpaceStartIndex)
				clearFile(fileBlock, length, originalStartIndex)
			}
			freeSpaceStartIndex = i
			freeSpaceCount = 1
			continue
		}

		if fileBlock[i-1] == -1 && fileBlock[i] != -1 {
			if freeSpaceCount >= length {
				writeFile(fileBlock, fileBlock[originalStartIndex], length, freeSpaceStartIndex)
				clearFile(fileBlock, length, originalStartIndex)
			}
			freeSpaceStartIndex = -1
			freeSpaceCount = 0
		}

		if fileBlock[i] == -1 {
			freeSpaceCount++
		}
	}
}

func writeFile(block []int, fileNum int, fileLength int, index int) {
	for i := range fileLength {
		block[index+i] = fileNum
	}
}

func clearFile(block []int, length int, index int) {
	for i := range length {
		block[index+i] = -1
	}
}

func calculateChecksum(diskBlocks []int) int {
	var counter = 0
	for i, block := range diskBlocks {
		if block != -1 {
			counter += i * block
		}
	}
	return counter
}
