package part2

import (
	"slices"
)

type RuleSet map[int][]int

func Solve(ruleSet RuleSet, pages [][]int) int {
	result := 0

	for _, page := range pages {
		if !isOrderValid(ruleSet, page) {
			fixPages := fixPageOrder(page, ruleSet)
			for {
				if isOrderValid(ruleSet, fixPages) {
					break
				}

				fixPages = fixPageOrder(fixPages, ruleSet)
			}

			middle := int(len(fixPages) / 2)
			result += fixPages[middle]
		}
	}

	return result
}

func fixPageOrder(pages []int, ruleSet RuleSet) []int {
	results := []struct {
		page int
		rule int
	}{}

	for i := len(pages) - 1; i >= 0; i-- {
		r := ruleSet[pages[i]]

		for s := i; s >= 0; s-- {
			for _, g := range r {
				if g == pages[s] {
					results = append(results, struct {
						page int
						rule int
					}{pages[i], g})
				}
			}
		}
	}

	slices.Reverse(results)

	for _, v := range results {
		index := indexOf(pages, v.page)
		pages = remove(pages, indexOf(pages, v.rule))
		pages = slices.Insert(pages, index, v.rule)

	}

	return pages
}

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
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
