package part2

import (
	"strings"
)

func Solve(input string) int {
	width := getWidth(input)
	input = strings.ReplaceAll(input, "\n", "")
	height := len(input) / width
	return searchForXMAS(input, width, height)
}

func searchForXMAS(input string, width, height int) int {
  _ = width
  _ = height

  result := 0
  for y := 1; y < height - 1 ; y++ {
    for x := 1; x < width - 1; x++ {
      r := ""
      r += string(input[(x+1)+width*(y-1)])
      r += string(input[(x+0)+width*(y+0)])
      r += string(input[(x-1)+width*(y+1)])

      l := ""
      l += string(input[(x+1)+width*(y+1)])
      l += string(input[(x+0)+width*(y+0)])
      l += string(input[(x-1)+width*(y-1)])

      if (l == "MAS" || l == "SAM") && (r == "MAS" || r == "SAM") {
        result++
      }
    }
  }

	return result
}

func getWidth(input string) int {
	for i := range input {
		if input[i] == '\n' {
			return i
		}
	}

	return len(input)
}
