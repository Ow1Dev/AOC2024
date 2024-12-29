package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Ow1Dev/AOC2024/day_5/part1"
	"github.com/Ow1Dev/AOC2024/day_5/part2"
)

func main() {
	content, err := os.ReadFile("./day_5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ruleSet, pages := getInput(string(content))

	fmt.Println("part1 result: ", part1.Solve(ruleSet, pages))
	fmt.Println("part2 result: ", part2.Solve(ruleSet, pages))
}

func getInput(content string) (map[int][]int, [][]int) {
	page_info_slice := strings.Split(strings.TrimSpace(content), "\n\n")

	rulesSet := make(map[int][]int)
	for _, v := range strings.Split(page_info_slice[0], "\n") {
		rules := strings.Split(v, "|")
		f, e := strconv.Atoi(rules[0])
		if e != nil {
			log.Fatal(e)
		}

		l, e := strconv.Atoi(rules[1])
		if e != nil {
			log.Fatal(e)
		}

		rulesSet[f] = append(rulesSet[f], l)
	}

	pages := make([][]int, 0)
	for _, v := range strings.Split(page_info_slice[1], "\n") {
		page := make([]int, 0)
		for _, p := range strings.Split(v, ",") {
			pa, e := strconv.Atoi(p)
			if e != nil {
				log.Fatal(e)
			}

			page = append(page, pa)
		}

		pages = append(pages, page)
	}

	return rulesSet, pages
}
