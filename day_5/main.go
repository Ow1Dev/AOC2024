package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./day_5/input.example.txt")
	if err != nil {
		log.Fatal(err)
	}

  page_info_str := string(content)
  page_info_slice := strings.Split(strings.TrimSpace(page_info_str), "\n\n")

  fmt.Println(page_info_slice[1])


	//fmt.Println("part1 result: ", part1.Solve(string(content)))
}
