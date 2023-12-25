package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err, "Failed to read input file")

	regex, err := regexp.Compile("[0-9]+")
	check(err, "Failed to process regex (you absolute idiot)")
	scanner := bufio.NewScanner(file)

	nextLine := nextLineFromScanner(scanner) // time
	time := regex.FindAllString(nextLine, -1)

	nextLine = nextLineFromScanner(scanner) // distance
	distance := regex.FindAllString(nextLine, -1)

	var differentOptions int = 1
	for i := 0; i < len(time); i++ {
		raceTime, err := strconv.Atoi(time[i])
		check(err, "Error parsing race time "+time[i])

		bestDistance, err := strconv.Atoi(distance[i])
		check(err, "Error parsing distance time "+distance[i])

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

		differentOptions *= highestBest - lowestBest + 1
	}

	fmt.Printf("The number of different ways to win is %d\n", differentOptions)
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
