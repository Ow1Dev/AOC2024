package part1

import (
	"strings"
)

func Solve(input string) int {
	width := getWidth(input)
	input = strings.ReplaceAll(input, "\n", "")
	height := len(input) / width

	resultV := searchV(input, width, height)
	resultH := searchH(input, width, height)
	resultD := searchD(input, width, height)

	return resultV + resultH + resultD
}

func searchD(input string, width, height int) int {
	result := 0

	for y := 0; y < height-3; y++ {
		for x := 0; x < width-3; x++ {
			r := string(input[x+width*y])
			r += string(input[(x+1)+width*(y+1)])
			r += string(input[(x+2)+width*(y+2)])
			r += string(input[(x+3)+width*(y+3)])

      if r == "XMAS" || r == "SAMX" {
        result++ 
      }
		}

		for x := 3; x < width; x++ {

      l := string(input[x+width*y])
			l += string(input[(x-1)+width*(y+1)])
			l += string(input[(x-2)+width*(y+2)])
			l += string(input[(x-3)+width*(y+3)])

      if l == "XMAS" || l == "SAMX" {
        result++ 
      }
		}
	}

	return result
}

func searchH(input string, width, height int) int {
	var result int
	for x := range width {
		for y := range height - 3 {
      s := string(input[x+width*(y+0)])
			s += string(input[(x)+width*(y+1)])
			s += string(input[(x)+width*(y+2)])
			s += string(input[(x)+width*(y+3)])

      if s == "XMAS" || s == "SAMX" {
        result++ 
      }
		}
	}
	return result
}

func searchV(input string, width, height int) int {
	var result int
	for x := range width - 3 {
		for y := range height {
      s := string(input[(x+0)+width*(y)])
			s += string(input[(x+1)+width*(y)])
			s += string(input[(x+2)+width*(y)])
			s += string(input[(x+3)+width*(y)])

      if s == "XMAS" || s == "SAMX" {
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

func check(code string, expect string) bool {
	charLen := len(expect)
	return len(code) >= charLen && code[0:charLen] == expect
}
