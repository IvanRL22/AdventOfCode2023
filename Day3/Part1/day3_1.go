package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPartNumber(data []string, line uint64, numberStart uint64, numberEnd uint64) bool {

	// Check previous line if not first line
	if line > 0 {
		if checkLine(data[line-1], numberStart, numberEnd) {
			return true
		}
	}

	// Check next line if not last line
	if line < uint64(len(data)-1) {
		if checkLine(data[line+1], numberStart, numberEnd) {
			return true
		}
	}

	if numberStart > 0 {
		numberStart--
	}
	if numberEnd < uint64(len(data[line]))-1 {
		numberEnd++
	}

	return checkChar(data[line][numberStart]) || checkChar(data[line][numberEnd])
}

func checkLine(line string, start uint64, end uint64) bool {
	if start > 0 {
		start--
	}
	if end < uint64(len(line)-1) {
		end++
	}

	for i := start; i <= end; i++ {
		if checkChar(line[i]) {
			return true
		}
	}

	return checkChar(line[start]) || checkChar(line[end])
}

func checkChar(char byte) bool {
	return char != '.' && (char < '0' || char > '9')
}

func main() {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	var total uint64
	for i := 0; i < len(data); i++ {
		dataLine := data[i]

		var firstDigit int = -1
		for j := 0; j < len(dataLine); j++ {
			if dataLine[j] >= '0' && dataLine[j] <= '9' {
				if firstDigit == -1 {
					firstDigit = j
				} else {
					for j < len(dataLine) {
						if dataLine[j] < '0' || dataLine[j] > '9' {
							break
						}
						j++
					}
					// TODO: encapsulate so it's not duplicated
					numberToCheck := dataLine[firstDigit:j]
					if isPartNumber(data, uint64(i), uint64(firstDigit), uint64(j-1)) {
						number, err := strconv.ParseUint(numberToCheck, 10, 32)
						check(err)

						total += number
					}
					firstDigit = -1
				}
			} else if firstDigit != -1 {
				// TODO: encapsulate so it's not duplicated
				numberToCheck := dataLine[firstDigit:j]
				if isPartNumber(data, uint64(i), uint64(firstDigit), uint64(j-1)) {
					number, err := strconv.ParseUint(numberToCheck, 10, 32)
					check(err)

					total += number
				}
				firstDigit = -1
			}

		}
	}

	fmt.Printf("The sum of all part numbers is %d", total)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
