package part1

import (
	"testing"

	"github.com/Ow1Dev/AOC2024/util"
)

func TestSolveSolution(t *testing.T) {
  rules := []Rule{
    {a: 47, b: 53},
    {a: 97, b: 13},
    {a: 97, b: 61},
    {a: 97, b: 47},
    {a: 75, b: 29},
    {a: 61, b: 13},
    {a: 75, b: 53},
    {a: 29, b: 13},
    {a: 97, b: 29},
    {a: 53, b: 29},
    {a: 61, b: 53},
    {a: 97, b: 53},
    {a: 61, b: 29},
    {a: 47, b: 13},
    {a: 75, b: 47},
    {a: 97, b: 75},
    {a: 47, b: 61},
    {a: 75, b: 61},
    {a: 47, b: 29},
    {a: 75, b: 13},
    {a: 53, b: 13},
  }
  
  pages := [][]int {
    {75,47,61,53,29},
    {97,61,53,29,13},
    {75,29,13},
    {75,97,47,61,53},
    {61,13,29},
    {97,13,75,29,47},
  }

	anwser := Solve(rules, pages)
	util.AssertEqual(t, anwser, 143, "value should match")

}
