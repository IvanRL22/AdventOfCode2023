package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err, "Failed to read input file")

	regex, err := regexp.Compile("[A-Z]{3}")
	check(err, "Failed to process regex (you absolute idiot)")
	scanner := bufio.NewScanner(file)

	instructions := nextLineFromScanner(scanner)
	nodes := make(map[string][2]string)

	nextLineFromScanner(scanner) //empty line

	for scanner.Scan() {
		node := regex.FindAllString(scanner.Text(), -1)
		nodes[node[0]] = [2]string{node[1], node[2]}
	}

	currentNode := nodes["AAA"]
	var steps uint = 0
	var endFound bool = false

	for endFound == false {
		for _, d := range instructions {
			steps++
			var nextNodeName string
			if d == 'L' {
				nextNodeName = currentNode[0]
			} else {
				nextNodeName = currentNode[1]
			}

			if nextNodeName == "ZZZ" {
				endFound = true
				break
			} else {
				currentNode = nodes[nextNodeName]
			}
		}
	}

	fmt.Printf("Total steps were %d\n", steps)
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
