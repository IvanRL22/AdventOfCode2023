package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input  string
		output int
	}{
		{
			// ???.### 1,1,3
			// .??..??...?##. 1,1,3
			// ?#?#?#?#?#?#?#? 1,3,1,6
			// ????.#...#... 4,1,1
			// ????.######..#####. 1,6,5
			// ?###???????? 3,2,1

			"???.### 1,1,3\n.??..??...?## 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1",
			21},
	}

	for i, ti := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			result := part1(ti.input)
			if result != ti.output {
				t.Errorf("Result %d when expected %d", result, ti.output)
			}
		})
	}
}

func TestCheckSpringsMatch(t *testing.T) {
	var tests = []struct {
		springs  string
		grouping []int
		output   bool
	}{
		{
			"#.#.###",
			[]int{1, 1, 3},
			true},
		{
			".#...#....###.",
			[]int{1, 1, 3},
			true},
		{
			".#.###.#.######",
			[]int{1, 3, 1, 6},
			true},
		{
			"#.#.#.#.###",
			[]int{1, 1, 1, 1, 2},
			false},
	}

	for i, ti := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			result := checkSpringsMatch(ti.springs, ti.grouping)
			if result != ti.output {
				t.Errorf("Result %t when expected %t", result, ti.output)
			}
		})
	}
}
