package part1

import (
	"slices"
)

func Solve(list1 []int, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)
	var result int

	for index := range list1 {
		var s int

		if list2[index] > list1[index] {
			s = list2[index] - list1[index]
		} else {
			s = list1[index] - list2[index]
		}

		result = result + s
	}

	return result
}
