package part1

import "testing"

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
  assertEqual(t, anwser, 18, "value should match")
}

func assertEqual(t *testing.T, got, want int, message string) {
	if got != want {
		t.Errorf("%s: got %d, expected %d", message, got, want)
	}
}
