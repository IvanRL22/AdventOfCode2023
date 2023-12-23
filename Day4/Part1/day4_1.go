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
	var totalScore uint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
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
		var cardScore uint
		numbersString := regex.FindAllString(barSplit[0], -1)
		for _, numS := range numbersString {
			number, err := strconv.Atoi(numS)
			check(err, "Failed to parse card number "+numS)

			if winners[number] != 0 {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}

		totalScore += cardScore
	}

	fmt.Printf("Total score is %d\n", totalScore)
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
