package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const cards string = "23456789TJQKA"

type Hand struct {
	cards string
	bid   int
	rank  int
}

// 5 > 4 > 3/2 > 3 > 2/2 > 2 > 1
// 7 > 6 >  5  > 4 >  3  > 2 > 1
func (h *Hand) getRank() int {
	handMap := make(map[rune]int)

	for _, c := range h.cards {
		handMap[c] += 1
	}

	var checkAgain int
	for _, v := range handMap {
		switch v {
		case 5:
			return 7
		case 4:
			return 6
		case 3, 2:
			if checkAgain == 0 {
				checkAgain = v
			} else if checkAgain == 3 { // v can only be 2, we have a full
				return 5
			} else if checkAgain == 2 {
				if v == 2 { // two 2s means double pair
					return 3
				} else {
					return 5 // 2 & 3 means a full
				}
			}
		}
	}

	// We have a set or a pair
	if checkAgain == 3 {
		return 4
	} else if checkAgain == 2 {
		return 2
	}

	return 1
}

func (a *Hand) lessByCards(b *Hand) bool {
	for i := 0; i < len(a.cards); i++ {
		indexA := strings.IndexByte(cards, a.cards[i])
		indexB := strings.IndexByte(cards, b.cards[i])
		if indexA != indexB {
			if indexA > indexB {
				return false
			} else {
				return true
			}
		}
	}

	return false // they're equal
}

func (a *Hand) compareByCards(b *Hand) int {
	for i := 0; i < len(a.cards); i++ {
		indexA := strings.IndexByte(cards, a.cards[i])
		indexB := strings.IndexByte(cards, b.cards[i])
		if indexA != indexB {
			if indexA > indexB {
				return 1
			} else {
				return -1
			}
		}
	}

	return 0
}

func main() {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	// file, err := os.Open("../reddit_data.txt")
	check(err, "Failed to read input file")
	scanner := bufio.NewScanner(file)

	hands := make([]*Hand, 0)

	for scanner.Scan() {
		handData := strings.Split(scanner.Text(), " ")

		bid, err := strconv.Atoi(handData[1])
		check(err, "Error parsing bid "+handData[1])

		hand := Hand{cards: handData[0], bid: bid}

		hands = append(hands, &hand)
	}

	sort.Slice(hands,
		func(i, j int) bool {
			byRank := hands[i].getRank() - hands[j].getRank()
			if byRank < 0 {
				return true
			} else if byRank > 0 {
				return false
			} else {
				return hands[i].lessByCards(hands[j])
			}
		})

	var totalPoints int
	for i, h := range hands {
		totalPoints += h.bid * (i + 1)
	}

	fmt.Printf("Total points are %d\n", totalPoints)
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
