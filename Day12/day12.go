package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("data.txt")
	check(err, "Failed to read input file")

	// fmt.Printf("Result of day 10 part 1 is %d\n", part1(string(file)))
	fmt.Printf("Result of day 10 part 2 is %d\n", part2(string(file)))
}

func part1(data string) int {
	combinationsMap := make(map[int][]string)

	var totalCombinations int
	for _, rowString := range strings.Split(data, "\n") {
		temp := strings.Split(rowString, " ")
		springs, groupingsS := temp[0], strings.Split(temp[1], ",")
		groupings := make([]int, len(groupingsS))

		for i, g := range groupingsS {
			groupings[i], _ = strconv.Atoi(g) // Is it ok to ignore err?
		}

		unknowns := strings.Count(springs, "?")
		var combinations []string
		if len(combinationsMap[unknowns]) == 0 {
			combinationsMap[unknowns] = generateCombinations(unknowns)
		}
		combinations = combinationsMap[unknowns]

		var goodCombinations int
		for _, comb := range combinations {
			test := springs
			for _, c := range comb {
				test = strings.Replace(test, "?", string(c), 1)
			}

			if checkSpringsMatch(test, groupings) {
				goodCombinations++
			}

		}
		totalCombinations += goodCombinations
	}

	return totalCombinations
}

func part2(data string) int {
	combinationsMap := make(map[int][]string)

	var totalCombinations int
	for _, rowString := range strings.Split(data, "\n") {
		temp := strings.Split(rowString, " ")
		springs, groupingsS := temp[0], strings.Split(temp[1], ",")
		groupings := make([]int, len(groupingsS))

		for i, g := range groupingsS {
			groupings[i], _ = strconv.Atoi(g) // Is it ok to ignore err?
		}

		// Part 2 specific
		springs = springs + springs + springs + springs + springs
		newGroupings := append(groupings, groupings...)
		newGroupings = append(newGroupings, groupings...)
		newGroupings = append(newGroupings, groupings...)
		newGroupings = append(newGroupings, groupings...)
		groupings = newGroupings

		unknowns := strings.Count(springs, "?")
		var combinations []string
		if len(combinationsMap[unknowns]) == 0 {
			combinationsMap[unknowns] = generateCombinations(unknowns)
		}
		combinations = combinationsMap[unknowns]

		var goodCombinations int
		for _, comb := range combinations {
			test := springs
			for _, c := range comb {
				test = strings.Replace(test, "?", string(c), 1)
			}

			if checkSpringsMatch(test, groupings) {
				goodCombinations++
			}

		}
		totalCombinations += goodCombinations
	}

	return totalCombinations
}

func generateCombinations(n int) []string {
	if n == 0 {
		return []string{""}
	}

	smaller := generateCombinations(n - 1)
	var result []string

	for _, s := range smaller {
		result = append(result, "."+s, "#"+s)
	}

	return result
}

func checkSpringsMatch(springs string, groupings []int) bool {
	regex, _ := regexp.Compile("#+")
	groupedSprings := regex.FindAllString(springs, -1)

	if len(groupedSprings) != len(groupings) {
		return false
	}

	for i, g := range groupedSprings {
		if len(g) != groupings[i] {
			return false
		}
	}

	return true
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
