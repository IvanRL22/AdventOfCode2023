package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.ReadFile("data.txt")
	check(err, "Failed to read input file")

	fmt.Printf("Result of day 10 part 1 is %d\n", part1(string(file)))
	fmt.Printf("Result of day 10 part 2 is %d\n", part2(string(file)))
}

func part1(data string) int {
	galaxies := make([][2]int, 0)
	emptyRows := make([]int, 0)

	dataM := strings.Split(data, "\n")
	for i, rowString := range dataM {
		currentGalaxyPos := strings.IndexRune(rowString, '#')
		if currentGalaxyPos == -1 {
			emptyRows = append(emptyRows, i)
			continue
		} else {
			galaxies = append(galaxies, [2]int{i, currentGalaxyPos})
			lastGalaxyPos := currentGalaxyPos + 1
			for {
				currentGalaxyPos := strings.IndexRune(rowString[lastGalaxyPos:], '#')
				if currentGalaxyPos == -1 {
					break
				} else {
					galaxies = append(galaxies, [2]int{i, currentGalaxyPos + lastGalaxyPos})
					lastGalaxyPos += currentGalaxyPos + 1
				}
			}
		}
	}

	emptyColumns := make([]int, 0)
	for j := 0; j < len(dataM[0]); j++ {
		for i := range dataM {
			if dataM[i][j] == '#' {
				break
			}
			if i == len(dataM)-1 {
				emptyColumns = append(emptyColumns, j)
			}
		}
	}

	var totalDistance int
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			startRow, startCol := galaxies[i][0], galaxies[i][1]
			endRow, endCol := galaxies[j][0], galaxies[j][1]

			distance := AbsInt(startRow-endRow) + AbsInt(startCol-endCol)

			for i := MinInt(startRow, endRow) + 1; i < MaxInt(startRow, endRow); i++ {
				if slices.Contains(emptyRows, i) {
					distance++
				}
			}
			for i := MinInt(startCol, endCol) + 1; i < MaxInt(startCol, endCol); i++ {
				if slices.Contains(emptyColumns, i) {
					distance++
				}
			}
			totalDistance += distance
		}
	}

	return totalDistance
}

// TODO Fix
// copy/pasted, would be better to decompose logic into functions and add parameter for expansion
func part2(data string) int {
	galaxies := make([][2]int, 0)
	emptyRows := make([]int, 0)

	dataM := strings.Split(data, "\n")
	for i, rowString := range dataM {
		currentGalaxyPos := strings.IndexRune(rowString, '#')
		if currentGalaxyPos == -1 {
			emptyRows = append(emptyRows, i)
			continue
		} else {
			galaxies = append(galaxies, [2]int{i, currentGalaxyPos})
			lastGalaxyPos := currentGalaxyPos + 1
			for {
				currentGalaxyPos := strings.IndexRune(rowString[lastGalaxyPos:], '#')
				if currentGalaxyPos == -1 {
					break
				} else {
					galaxies = append(galaxies, [2]int{i, currentGalaxyPos + lastGalaxyPos})
					lastGalaxyPos += currentGalaxyPos + 1
				}
			}
		}
	}

	emptyColumns := make([]int, 0)
	for j := 0; j < len(dataM[0]); j++ {
		for i := range dataM {
			if dataM[i][j] == '#' {
				break
			}
			if i == len(dataM)-1 {
				emptyColumns = append(emptyColumns, j)
			}
		}
	}

	var totalDistance int
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			startRow, startCol := galaxies[i][0], galaxies[i][1]
			endRow, endCol := galaxies[j][0], galaxies[j][1]

			distance := AbsInt(startRow-endRow) + AbsInt(startCol-endCol)

			for i := MinInt(startRow, endRow) + 1; i < MaxInt(startRow, endRow); i++ {
				if slices.Contains(emptyRows, i) {
					distance += 999999
				}
			}
			for i := MinInt(startCol, endCol) + 1; i < MaxInt(startCol, endCol); i++ {
				if slices.Contains(emptyColumns, i) {
					distance += 999999
				}
			}
			totalDistance += distance
		}
	}

	return totalDistance
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func check(e error, message string) {
	if e != nil {
		panic(message)
	}
}
