package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// file, err := os.Open("example.txt")
	file, err := os.Open("data.txt")
	check(err)

	var total int
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		var first, last int
		for _, currentChar := range line {
			if currentChar > '0' && currentChar <= '9' {
				currentNumber, err := strconv.Atoi(string(currentChar))
				check(err)

				if first == 0 {
					first = currentNumber
				} else {
					last = currentNumber
				}
			}
		}

		if last == 0 {
			last = first
		}

		lineNumber := first*10 + last
		fmt.Printf("First number is %d, last number is %d for a line number of is %d\n", first, last, lineNumber)

		total += lineNumber
	}

	fmt.Printf("Total is %d", total)

	file.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
