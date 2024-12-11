package part2

import (
	"testing"

	"github.com/Ow1Dev/AOC2024/util"
)

func TestSolveSolution(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

	anwser := Solve(input)
	util.AssertEqual(t, anwser, 9, "value should match")
}
