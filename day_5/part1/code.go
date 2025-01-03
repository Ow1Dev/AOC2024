package part1

func Solve(ruleSet map[int][]int, pages [][]int) int {
	result := 0

	for _, page := range pages {
		if isOrderValid(ruleSet, page) {
			middle := int(len(page) / 2)
			result += page[middle]
		}
	}

	return result
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
