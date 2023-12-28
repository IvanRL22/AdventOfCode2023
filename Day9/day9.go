package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("data.txt")
	check(err, "Failed to read input file")

	part1(file)
	part2(file)

}

func part1(data []byte) {
	var finalResult int

	for _, rowString := range strings.Split(strings.TrimSpace(string(data)), "\r\n") {
		numbersString := strings.Split(rowString, " ")
		currentNumbers := make([]int, len(numbersString))
		calcNumbers := make([]int, 0)

		for i, s := range numbersString {
			n, err := strconv.Atoi(s)
			check(err, "Error parsing number "+s)

			currentNumbers[i] = n
		}

		endNumber := currentNumbers[len(currentNumbers)-1]

		for slices.Max(currentNumbers) != 0 || slices.Min(currentNumbers) != 0 {
			nextNumbers := make([]int, len(currentNumbers)-1)

			for i := 0; i < len(currentNumbers)-1; i++ {
				nextNumbers[i] = currentNumbers[i+1] - currentNumbers[i]
			}

			calcNumbers = append(calcNumbers, nextNumbers[len(nextNumbers)-1])
			currentNumbers = nextNumbers
		}

		var finalNumber int
		for _, n := range calcNumbers {
			finalNumber += n
		}

		finalResult += finalNumber + endNumber // Need to add the last number from the original row
	}

	fmt.Printf("Part 1 result is %d\n", finalResult)
}

func part2(data []byte) {
	var finalResult int

	for _, rowString := range strings.Split(strings.TrimSpace(string(data)), "\r\n") {
		numbersString := strings.Split(rowString, " ")
		currentNumbers := make([]int, len(numbersString))
		calcNumbers := make([]int, 0)

		for i, s := range numbersString {
			n, err := strconv.Atoi(s)
			check(err, "Error parsing number "+s)

			currentNumbers[i] = n
		}

		startNumber := currentNumbers[0]

		for slices.Max(currentNumbers) != 0 || slices.Min(currentNumbers) != 0 {
			nextNumbers := make([]int, len(currentNumbers)-1)

			for i := 0; i < len(currentNumbers)-1; i++ {
				nextNumbers[i] = currentNumbers[i+1] - currentNumbers[i]
			}

			calcNumbers = append(calcNumbers, nextNumbers[0])
			currentNumbers = nextNumbers
		}

		var finalNumber int
		for i := len(calcNumbers) - 1; i >= 0; i-- {
			finalNumber = calcNumbers[i] - finalNumber
		}

		finalResult += startNumber - finalNumber // Need to calculate the first number from the original row
	}

	fmt.Printf("Part 2 result is %d\n", finalResult)
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
