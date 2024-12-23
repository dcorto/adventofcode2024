package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

const day = 23

// lan represents a LAN connection of 3 computers
type lan struct {
	a, b, c string
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
	input, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	lanMap := createMap(input)
	solution = findConnectionsWithT(findThreeConnectedComputers(lanMap))
	return solution
}

func solutionB() string {
	var solution = ""
	input, _ := utils.ReadLinesFromFile(fmt.Sprintf("%d/input.txt", day))
	lanMap := createMap(input)
	solution = findMaxCliques(lanMap)
	return solution
}

// createMap creates a map of maps to represent the LAN connections
func createMap(input []string) map[string]map[string]struct{} {
	lanMap := make(map[string]map[string]struct{})

	for _, line := range input {
		computer1, computer2 := line[:2], line[3:]

		if _, ok := lanMap[computer1]; !ok {
			lanMap[computer1] = make(map[string]struct{})
		}

		if _, ok := lanMap[computer2]; !ok {
			lanMap[computer2] = make(map[string]struct{})
		}

		lanMap[computer1][computer2] = struct{}{}
		lanMap[computer2][computer1] = struct{}{}
	}
	return lanMap
}

// findThreeConnectedComputers finds all 3 connected computers
func findThreeConnectedComputers(lanMap map[string]map[string]struct{}) map[lan]struct{} {
	result := make(map[lan]struct{})
	for key, value := range lanMap {
		if len(value) <= 1 {
			continue
		}

		for key2 := range value {
			for key3 := range value {
				if key2 == key3 {
					continue
				}

				if _, ok := lanMap[key2][key3]; ok {
					r := []string{key, key2, key3}
					sort.Strings(r)
					result[lan{r[0], r[1], r[2]}] = struct{}{}
				}
			}
		}
	}
	return result
}

// findConnectionsWithT finds all connections with a computer starting with "t"
func findConnectionsWithT(input map[lan]struct{}) int {
	count := 0
	for key := range input {
		if strings.HasPrefix(key.a, "t") || strings.HasPrefix(key.b, "t") || strings.HasPrefix(key.c, "t") {
			count++
		}
	}
	return count
}

// BronKerbosch is an implementation of the Bron-Kerbosch algorithm to find all maximal cliques in a graph
func BronKerbosch(currentClique []string, yetToConsider []string, alreadyConsidered []string, lanMap map[string]map[string]struct{}, cliques [][]string) [][]string {
	if len(yetToConsider) == 0 && len(alreadyConsidered) == 0 {
		cliques = append(cliques, append([]string{}, currentClique...))
		return cliques
	}

	for index := 0; index < len(yetToConsider); {
		node := yetToConsider[index]
		var newYetToConsider []string
		var newAlreadyConsidered []string

		for _, n := range yetToConsider {
			if _, ok := lanMap[node][n]; ok {
				newYetToConsider = append(newYetToConsider, n)
			}
		}

		for _, n := range alreadyConsidered {
			if _, ok := lanMap[node][n]; ok {
				newAlreadyConsidered = append(newAlreadyConsidered, n)
			}
		}

		cliques = BronKerbosch(append(currentClique, node), newYetToConsider, newAlreadyConsidered, lanMap, cliques)

		yetToConsider = append(yetToConsider[:index], yetToConsider[index+1:]...)
		alreadyConsidered = append(alreadyConsidered, node)
	}
	return cliques
}

// findMaxCliques finds the maximal cliques in the LAN map
func findMaxCliques(lanMap map[string]map[string]struct{}) string {
	var maxClique []string
	var allComputers []string

	for key := range lanMap {
		allComputers = append(allComputers, key)
	}

	cliques := BronKerbosch([]string{}, allComputers, []string{}, lanMap, [][]string{})

	for _, c := range cliques {
		if len(c) > len(maxClique) {
			maxClique = c
		}
	}

	sort.Strings(maxClique)
	return strings.Join(maxClique, ",")
}
