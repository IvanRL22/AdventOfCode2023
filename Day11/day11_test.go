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
			// ...#......
			// .......#..
			// #.........
			// ..........
			// ......#...
			// .#........
			// .........#
			// ..........
			// .......#..
			// #...#.....
			"...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....",
			374},
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
