package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

// Function to find gcd of two numbers
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Function to find lcm of two numbers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Function to find lcm of an array of numbers
func lcmArray(arr []int) int {
	num1 := arr[0]
	num2 := arr[1]
	temp := lcm(num1, num2)

	for i := 2; i < len(arr); i++ {
		temp = lcm(temp, arr[i])
	}

	return temp
}

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("../data.txt")
	check(err, "Failed to read input file")

	regex, err := regexp.Compile("[0-9A-Z]{3}")
	check(err, "Failed to process regex (you absolute idiot)")
	scanner := bufio.NewScanner(file)

	instructions := nextLineFromScanner(scanner)
	nodes := make(map[string][2]string)
	currentNodes := make([]string, 0)

	nextLineFromScanner(scanner) //empty line

	for scanner.Scan() {
		node := regex.FindAllString(scanner.Text(), -1)
		nodes[node[0]] = [2]string{node[1], node[2]}

		if node[0][2] == 'A' {
			currentNodes = append(currentNodes, node[0])
		}
	}

	loops := make([]int, len(currentNodes))

	var currentSteps int = 0
	var endFound bool = false
	for endFound == false {
		for _, d := range instructions {

			var nextStep int
			if d == 'L' {
				nextStep = 0
			} else {
				nextStep = 1
			}
			currentSteps++

			for i, current := range currentNodes {
				if loops[i] == 0 {
					currentNodes[i] = nodes[current][nextStep]
					if currentNodes[i][2] == 'Z' {
						loops[i] = currentSteps
					}
				}
			}

			if slices.Min(loops) != 0 {
				endFound = true
				break
			}
		}
	}

	fmt.Printf("The total number of steps is %d\n", lcmArray(loops))
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
