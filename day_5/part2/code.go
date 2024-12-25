package part2

import "fmt"

type RuleSet map[int][]int

func Solve(ruleSet RuleSet, pages [][]int) int {
	result := 0

	for _, page := range pages {
		if !isOrderValid(ruleSet, page) {
      fixPages := fixPageOrder(page, ruleSet)
      fmt.Println(fixPages)
		}
	}

	return result
}

func fixPageOrder(pages []int, ruleSet RuleSet) []int {
	for i := len(pages) - 1; i >= 0; i-- {
		r := ruleSet[pages[i]]

		for s := i; s >= 0; s-- {
			for _, g := range r {
				if g == pages[s] {
          // TODO: we need to reverse this order and then do the operations
          // first we need to get current page number and put it before current reverse this order and then do the operations
          fmt.Println("swp", pages[i], g)
				}
			}
		}
	}
  return pages
}

func isOrderValid(rules map[int][]int, pages []int) bool {
	for i := len(pages) - 1; i >= 0; i-- {
		r := rules[pages[i]]

		for s := i; s >= 0; s-- {
			for _, g := range r {
				if g == pages[s] {
					return false
				}
			}
		}
	}

	return true
}
