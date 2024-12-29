package part1

import (
	"testing"

	"github.com/Ow1Dev/AOC2024/util"
)

func TestSolveSolution(t *testing.T) {
  input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	anwser := Solve(input)
	util.AssertEqual(t, anwser, 41, "value should match")
}
