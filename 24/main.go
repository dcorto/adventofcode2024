package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

const day = 24

func main() {
	fmt.Println("Solution for Day", day)

	solutionA := solutionA()
	fmt.Println("Solution A:", solutionA)

	solutionB := solutionB()
	fmt.Println("Solution B:", solutionB)
}

func solutionA() int {

	var solution = 0
	input, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))

	initialValues, gates := parseInputForPartOne(input)

	// Store wire values
	wireValues := make(map[string]int)
	for k, v := range initialValues {
		wireValues[k] = v
	}

	// Get the value of a wire
	getWireValue := func(wire string) int {
		if val, ok := wireValues[wire]; ok {
			return val
		}
		return -1
	}

	// Process the gates
	loop := true
	for loop {
		loopAgain := false
		for _, gate := range gates {
			parts := strings.Split(gate, " ")
			input1, operation, input2, output := parts[0], parts[1], parts[2], parts[4]

			leftValue, rightValue := getWireValue(input1), getWireValue(input2)

			if leftValue == -1 || rightValue == -1 {
				loopAgain = true
				continue
			}

			var result int

			switch operation {
			case "AND":
				result = leftValue & rightValue
			case "OR":
				result = leftValue | rightValue
			case "XOR":
				result = leftValue ^ rightValue
			}

			wireValues[output] = result
		}
		loop = loopAgain
	}

	// Combine the values of wires starting with z
	var zKeys []string
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if val, ok := wireValues[wire]; ok {
			zKeys = append(zKeys, strconv.Itoa(val))
		} else {
			break
		}
	}

	// Reverse the order of the bits
	for i, j := 0, len(zKeys)-1; i < j; i, j = i+1, j-1 {
		zKeys[i], zKeys[j] = zKeys[j], zKeys[i]
	}

	// Convert the binary number to decimal
	var binaryResult strings.Builder
	for _, val := range zKeys {
		binaryResult.WriteString(val)
	}

	binaryString := binaryResult.String()
	decimalResult, _ := strconv.ParseInt(binaryString, 2, 64)

	solution = int(decimalResult)
	return solution
}

func solutionB() string {
	var solution = ""
	input, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))

	parts := parseInputForPartTwo(input)
	solution = swapAndJoinWires(parts)
	return solution
}

// parseInputForPartOne reads the input file and splits it into a map of initialValues and a slice of gates.
func parseInputForPartOne(input []string) (map[string]int, []string) {
	initialValues := make(map[string]int)
	var gates []string

	for _, line := range input {

		if line == "" {
			continue
		}

		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			parts = strings.Split(line, ": ")
			initialValues[parts[0]], _ = strconv.Atoi(parts[1])
			continue
		}

		value, output := parts[0], parts[1]
		if strings.Contains(value, "AND") || strings.Contains(value, "OR") || strings.Contains(value, "XOR") {
			gates = append(gates, value+" -> "+output)
		}
	}

	return initialValues, gates
}

// parseInputForPartTwo reads the input file and splits it into parts.
func parseInputForPartTwo(input []string) [][]string {
	var parts [][]string
	var currentPart []string

	for _, line := range input {
		if line == "" {
			if len(currentPart) > 0 {
				parts = append(parts, currentPart)
				currentPart = nil
			}
		} else {
			currentPart = append(currentPart, line)
		}
	}
	if len(currentPart) > 0 {
		parts = append(parts, currentPart)
	}

	return parts
}

// swapAndJoinWires performs the logic to swap and join wires.
func swapAndJoinWires(data [][]string) string {
	gates := data[1]
	var swapped []string
	var c0 string

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		// Half adder logic
		m1 = findGate("x"+n, "y"+n, "XOR", gates)
		n1 = findGate("x"+n, "y"+n, "AND", gates)

		if c0 != "" {
			r1 = findGate(c0, m1, "AND", gates)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = findGate(c0, m1, "AND", gates)
			}

			z1 = findGate(c0, m1, "XOR", gates)

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}

			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}

			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = findGate(r1, n1, "OR", gates)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if c0 == "" {
			c0 = n1
		} else {
			c0 = c1
		}
	}

	// Sort and join swapped wires
	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}

// findGate finds the gate that contains the given inputs and operator.
func findGate(a, b, operator string, gates []string) string {
	for _, gate := range gates {
		if strings.Contains(gate, fmt.Sprintf("%s %s %s", a, operator, b)) ||
			strings.Contains(gate, fmt.Sprintf("%s %s %s", b, operator, a)) {
			return strings.Split(gate, " -> ")[1]
		}
	}
	return ""
}
