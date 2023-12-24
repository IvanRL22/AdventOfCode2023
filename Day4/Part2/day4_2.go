package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err, "Failed to read input file")

	regex, err := regexp.Compile("[0-9]+")
	var totalScratchcardsWon int
	scratchcardsWon := make(map[int]int)

	var currentLine int = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Total number of card %d is %d\n", currentLine, scratchcardsWon[currentLine]+1)
		line := scanner.Text()

		colonSplit := strings.Split(line, ":")
		barSplit := strings.Split(colonSplit[1], "|")

		// Save winners into a map (kinda ugly but no sets)
		winners := make(map[int]int)
		for _, winS := range regex.FindAllString(barSplit[1], -1) {
			winner, err := strconv.Atoi(winS)
			check(err, "Failed to parse winner "+winS)
			winners[winner] = winner
		}

		// Check numbers against map
		var cardScore int
		numbersString := regex.FindAllString(barSplit[0], -1)
		for _, numS := range numbersString {
			number, err := strconv.Atoi(numS)
			check(err, "Failed to parse card number "+numS)

			if winners[number] != 0 {
				cardScore++
			}
		}

		// Increment amount of next scratchcards won (only the current card amount matters)
		for i := 1; i <= cardScore; i++ {
			scratchcardsWon[currentLine+i] += scratchcardsWon[currentLine] + 1
		}

		totalScratchcardsWon += scratchcardsWon[currentLine] + 1
		currentLine++
	}

	fmt.Printf("Total score is %d\n", totalScratchcardsWon)
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
