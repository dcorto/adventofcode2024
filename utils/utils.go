package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return lines, nil
}

func SliceFromStringToInt(slice []string) ([]int, error) {
	var intSlice []int
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}
