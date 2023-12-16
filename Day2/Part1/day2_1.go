package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id   uint
	sets []set
}

type set struct {
	red   uint
	green uint
	blue  uint
}

func gameFromString(line string) game {
	var tokens = strings.Split(line, ":")
	var gameId, err = strconv.Atoi(tokens[0][5:])
	check(err)

	newGame := game{id: uint(gameId)}
	newGame.fillSets(tokens[1])
	return newGame
}

func (g *game) fillSets(lineSets string) {
	var sets = strings.Split(lineSets, ";")
	finalGameSets := make([]set, len(sets))

	for i := 0; i < len(sets); i++ {
		newSet := set{}
		gameSet := strings.Split(sets[i], ",")

		for j := 0; j < len(gameSet); j++ {
			cubeSet := strings.Split(strings.Trim(gameSet[j], " "), " ")
			colorOfCube := cubeSet[1]
			numberOfCubes, err := strconv.Atoi(cubeSet[0])
			check(err)

			if strings.EqualFold(colorOfCube, "red") {
				newSet.red = uint(numberOfCubes)
			} else if strings.EqualFold(colorOfCube, "green") {
				newSet.green = uint(numberOfCubes)
			} else if strings.EqualFold(colorOfCube, "blue") {
				newSet.blue = uint(numberOfCubes)
			} else {
				panic("WTF?")
			}
		}

		finalGameSets[i] = newSet
	}

	g.sets = finalGameSets
}

func (s set) isPossibleWith(anotherSet set) bool {
	if s.red > anotherSet.red {
		return false
	}
	if s.green > anotherSet.green {
		return false
	}
	if s.blue > anotherSet.blue {
		return false
	}

	return true
}

func (g *game) isPossible(s set) bool {
	if len(g.sets) == 0 {
		return false
	}

	for i := 0; i < len(g.sets); i++ {
		if g.sets[i].isPossibleWith(s) {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	setToCompare := set{red: 12, green: 13, blue: 14}

	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err)

	var line string
	var total uint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		newGame := gameFromString(line)

		if newGame.isPossible(setToCompare) {
			total += newGame.id
		}
	}

	fmt.Printf("Total of sum of possible game IDs is %d\n", total)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
