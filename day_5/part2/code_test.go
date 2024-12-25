package part2

import (
	"testing"

	"github.com/Ow1Dev/AOC2024/util"
)

func TestSolveSolution(t *testing.T) {
	rulesSet := map[int][]int{
		47: {53, 13, 61, 29},
		97: {13, 61, 47, 29, 53, 75},
		75: {29, 53, 47, 61, 13},
		61: {13, 53, 29},
		29: {13},
		53: {29, 13},
	}

	pages := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	anwser := Solve(rulesSet, pages)
	util.AssertEqual(t, anwser, 123, "value should match")
}
