package part1

import "testing"

func TestSolveSolution(t *testing.T) {
  input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

  anwser := Solve(input)

  if anwser != 161 {
    t.Errorf("got %d, expected %d", anwser, 161)
  }
}
