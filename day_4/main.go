package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ow1Dev/AOC2024/day_4/part1"
)

func main() {
	content, err := os.ReadFile("./day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

  fmt.Println("part1 result: ", part1.Solve(string(content)))
}
