package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Ow1Dev/AOC2024/day_2/part1"
	"github.com/Ow1Dev/AOC2024/day_2/part2"
)

func main() {
  inputs, err := getInput("./day_2/input.txt")
  if err != nil {
    log.Fatal("Something went worng")
  }

  var result1, result2 int

  for _, input := range inputs {
    s := part1.Solve(input)
    if s == true {
      result1++
    }

    s = part2.Solve(input)
    if s == true {
      result2 = result2 + 1
    }
  }

  fmt.Println("part1 result: ", result1)
  fmt.Println("part2 result: ", result2)
}

func getInput(path string) ([][]int, error) {
	f, err := os.Open(path)
	if err != nil {
    return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

  list := make([][]int, 0)

	for scanner.Scan() {
    numbers := strings.Split(scanner.Text(), " ")
    line, err := convertStringsToInts(numbers)
    if err != nil {
      return nil, err
    }

    list = append(list, line)
  }

  return list, nil
} 

func convertStringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error converting %q to int: %v", s, err)
		}
		ints[i] = num
	}
	return ints, nil
}
