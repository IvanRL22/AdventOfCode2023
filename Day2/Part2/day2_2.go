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

func (g *game) findMinimumViableSet() *set {
	minimumViableSet := set{red: 0, green: 0, blue: 0}

	for i := 0; i < len(g.sets); i++ {
		if g.sets[i].red > minimumViableSet.red {
			minimumViableSet.red = g.sets[i].red
		}
		if g.sets[i].green > minimumViableSet.green {
			minimumViableSet.green = g.sets[i].green
		}
		if g.sets[i].blue > minimumViableSet.blue {
			minimumViableSet.blue = g.sets[i].blue
		}
	}

	return &minimumViableSet
}

func (s *set) getPower() uint {
	return s.red * s.green * s.blue
}

func main() {

	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	check(err)

	var line string
	var total uint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		newGame := gameFromString(line)
		total += newGame.findMinimumViableSet().getPower()
	}

	fmt.Printf("Total of sum of the power of all minimum viable sets is %d\n", total)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
