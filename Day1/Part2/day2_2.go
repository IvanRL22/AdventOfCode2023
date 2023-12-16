package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	// file, err := os.Open("example.txt")
	file, err := os.Open("../data.txt")
	check(err)

	numbersString := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbersChar := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var total int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var first, last int = 10, 0
		var firstPosition, lastPositon int = math.MaxInt, -1

		for i := 0; i < len(numbersString); i++ {

			var currentStringFirstIndex = strings.Index(line, numbersString[i])
			if currentStringFirstIndex != -1 {
				if currentStringFirstIndex < firstPosition {
					firstPosition = currentStringFirstIndex
					first = numbers[i]
				}

				var currentStringLastIndex = strings.LastIndex(line, numbersString[i])
				if currentStringLastIndex > lastPositon {
					lastPositon = currentStringLastIndex
					last = numbers[i]
				}
			}

			var currentDigitIndex = strings.Index(line, numbersChar[i])
			if currentDigitIndex != -1 {
				if currentDigitIndex < firstPosition {
					firstPosition = currentDigitIndex
					first = numbers[i]
				}

				var currentDigitLastIndex = strings.LastIndex(line, numbersChar[i])
				if currentDigitLastIndex > lastPositon {
					lastPositon = currentDigitLastIndex
					last = numbers[i]
				}
			}

		}

		var lineNumber = first*10 + last
		fmt.Printf("First number is %d, last number is %d for a line number of is %d\n", first, last, lineNumber)

		total += lineNumber
	}

	fmt.Printf("Total is %d\n", total)

	file.Close()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
