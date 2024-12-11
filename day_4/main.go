package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ow1Dev/AOC2024/day_4/part1"
	"github.com/Ow1Dev/AOC2024/day_4/part2"
)

func main() {
	content, err := os.ReadFile("./day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1 result: ", part1.Solve(string(content)))
	fmt.Println("part2 result: ", part2.Solve(string(content)))
}
