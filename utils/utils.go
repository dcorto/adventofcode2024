package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadLinesFromFile reads a file and returns a slice of strings one for each line
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

// ReadFromFile reads a file and returns it's content as a string
func ReadFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	content := string(data)
	return content, nil
}

// SliceFromStringToInt converts a slice of strings to a slice of integers
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

// SliceFromStringToInt64 converts a slice of strings to a slice of int64
func SliceFromStringToInt64(slice []string) ([]int64, error) {
	var intSlice []int64
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		intSlice = append(intSlice, int64(i))
	}
	return intSlice, nil
}
