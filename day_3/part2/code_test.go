package part2

import "testing"

func TestSolveSolution(t *testing.T) {
  input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
  expect := 48

  anwser := Solve(input)

  if anwser != expect {
    t.Errorf("got %d, expected %d", anwser, expect)
  }
}
