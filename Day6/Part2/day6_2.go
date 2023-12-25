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
	check(err, "Failed to process regex (you absolute idiot)")
	scanner := bufio.NewScanner(file)

	nextLine := nextLineFromScanner(scanner) // time
	timeString := strings.Join(regex.FindAllString(nextLine, -1), "")

	nextLine = nextLineFromScanner(scanner) // distance
	distanceString := strings.Join(regex.FindAllString(nextLine, -1), "")

	raceTime, err := strconv.Atoi(timeString)
	check(err, "Error parsing race time "+timeString)

	bestDistance, err := strconv.Atoi(distanceString)
	check(err, "Error parsing distance time "+distanceString)

	var lowestBest, highestBest = 0, raceTime
	// Find lowest time that wins
	for currentTime := 1; currentTime < raceTime; currentTime++ {
		if currentTime*(raceTime-currentTime) > bestDistance {
			lowestBest = currentTime
			break
		}
	}

	// Find highest time that wins
	for currentTime := raceTime - 1; currentTime > lowestBest; currentTime-- {
		if currentTime*(raceTime-currentTime) > bestDistance {
			highestBest = currentTime
			break
		}
	}

	fmt.Printf("The number of different ways to win is %d\n", highestBest-lowestBest+1)
}

func nextLineFromScanner(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
