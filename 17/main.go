package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"utils"
)

const day = 17

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() string {
	var solution = ""

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", 17))
	if err != nil {
		fmt.Println("Error:", err)
	}

	registerA, registerB, registerC, program := parseInput(lines)
	output := runProgram(registerA, registerB, registerC, program)
	solution = parseOutput(output)
	return solution
}

func solutionB() int {
	var solution = 0

	lines, err := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", 17))
	if err != nil {
		fmt.Println("Error:", err)
	}

	registerA, registerB, registerC, program := parseInput(lines)
	//output := runProgram(registerA, registerB, registerC, program)
	//solution = parseOutput(output)

	solution = int(findValueOfA(registerA, registerB, registerC, program))

	return solution
}

func getOperandValue(input int64, registerA, registerB, registerC int64) int64 {
	switch input {
	case 0, 1, 2, 3:
		return input
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	case 7:
		return 7
	}
	return input
}

func processOpcode(operand int64, opcode int64, registerA, registerB, registerC int64, ip int64) (int64, int64, int64, int64, int64) {
	operandCalculated := getOperandValue(operand, registerA, registerB, registerC)
	var output int64
	output = -1

	switch opcode {
	case 0:
		registerA >>= operandCalculated
	case 1:
		registerB ^= operand //7
	case 2:
		registerB = operandCalculated & 7
	case 3:
		if registerA != 0 {
			ip = operand
		}
	case 4:
		registerB ^= registerC
	case 5:
		output = operandCalculated & 7
	case 6:
		registerB = registerA >> operandCalculated
	case 7:
		registerC = registerA >> operandCalculated
	}
	return registerA, registerB, registerC, ip, output
}

func runProgram(a, b, c int64, inst []int64) []int64 {
	var ip int64
	var output []int64
	for ip < int64(len(inst)-1) {
		var newIP, o int64
		newIP = ip
		o = -1
		a, b, c, newIP, o = processOpcode(inst[ip+1], inst[ip+0], a, b, c, ip)
		if newIP == ip {
			ip += 2
		} else {
			ip = newIP
		}
		if o != -1 {
			output = append(output, o)
		}
	}
	return output
}

func findValueOfA(a, b, c int64, inst []int64) int64 {
	var output []int64
	a = 1

	for {
		output = runProgram(a, b, c, inst)

		if reflect.DeepEqual(output, inst) {
			return a
		}

		if len(inst) > len(output) {
			a *= 2
			continue
		}

		if len(inst) == len(output) {
			for j := len(inst) - 1; j >= 0; j-- {
				if inst[j] != output[j] {
					// Key Insight: every nth digit increments at every 8^n th step.
					// https://www.reddit.com/r/adventofcode/comments/1hg38ah/comment/m2gkd6m/
					a += int64(math.Pow(8, float64(j)))
					break
				}
			}
		}

		if len(inst) < len(output) {
			a /= 2
		}
	}
}

func parseInput(lines []string) (a, b, c int64, program []int64) {
	endRegisters := false

	for _, line := range lines {
		if !endRegisters {
			if len(line) == 0 {
				endRegisters = true
			} else {
				var value int64
				var register string
				fmt.Sscanf(line, "Register %1s: %d", &register, &value)
				switch register {
				case "A":
					a = value
					break
				case "B":
					b = value
					break
				case "C":
					c = value
					break

				}

			}
		} else {
			line = strings.TrimPrefix(line, "Program: ")
			program, _ = utils.SliceFromStringToInt64(strings.Split(line, ","))
		}
	}
	return a, b, c, program
}

func parseOutput(output []int64) string {
	strSlice := make([]string, len(output))
	for i, v := range output {
		strSlice[i] = strconv.Itoa(int(v))
	}
	return strings.Join(strSlice, ",")
}
