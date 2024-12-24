package part1

type Rule struct {
	a int
	b int
}


func Solve(rules []Rule, pages [][]int) int {
  result := 0

	rulesMap := make(map[int][]int)
	for _, v := range rules {
		rulesMap[v.a] = append(rulesMap[v.a], v.b)
	}

	for _, page := range pages {
		if isOrderValid(rulesMap, page) {
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

// 29 53 61 47 75
// ^
