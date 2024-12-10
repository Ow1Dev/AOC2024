package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Ow1Dev/AOC2024/day_1/part1"
	"github.com/Ow1Dev/AOC2024/day_1/part2"
)

func main() {
	list1, list2 := ReadInput("./day_1/input.txt")
  
  fmt.Println("part1 result: ", part1.Solve(list1, list2))
  fmt.Println("part2 result: ", part2.Solve(list1, list2))
}

func ReadInput(path string) ([]int, []int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	list1, list2 := make([]int, 0), make([]int, 0)

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), " ")
		number1, _ := strconv.Atoi(strings.TrimSpace(numbers[0]))
		number2, _ := strconv.Atoi(strings.TrimSpace(numbers[len(numbers)-1]))

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}
	f.Close()

	return list1, list2
}
