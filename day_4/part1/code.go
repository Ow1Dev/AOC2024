package part1

import (
	"strings"
)

func Solve(input string) int {
  result := 0

  width := getWidth(input)
  input = strings.ReplaceAll(input, "\n", "")
  height := len(input) / width 

  result += searchV(input) 
  result += searchH(input, width, height) 
  result += searchD(input, width, height) 
  return result 
}

// ...xxxx.....xxxxxx...xxxxxxxx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.xxxxxxxx...xxxxxx.....xxxx...
func searchD(input string, width, height int) int {
  var result int
  for x := range width {
    for y := range height - 3 {
      s := string(input[x + width * y + 1])
      s += string(input[x + width * y + 2])
      s += string(input[x + width * y + 3])
      s += string(input[x + width * y + 4])

      result += searchV(s) 
    }
  }
  return result 
}

func searchH(input string, width, height int) int {
  var result int
  for x := range width {
    for y := range height - 3 {
      s := string(input[x + width * y + 1])
      s += string(input[x + width * y + 2])
      s += string(input[x + width * y + 3])
      s += string(input[x + width * y + 4])

      result += searchV(s) 
    }
  }
  return result 
}

func searchV(input string) int {
  result := 0
  for {
    if check(input, "XMAS") {
      result++
    }

    if check(input, "SAMX") {
      result++
    }

		input = input[1:]
		if len(input) == 0 {
      break
		}
  }
  return result
}

func getWidth(input string) int {
  for i := range input{
    if (input[i] == '\n') {
      return i
    }
  }

  return len(input)
}

func check(code string, expect string) bool {
  charLen := len(expect)
  return len(code) >= charLen && code[0:charLen] == expect 
}
