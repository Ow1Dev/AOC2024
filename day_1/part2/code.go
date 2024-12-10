package part2

func Solve(list1 []int, list2 []int) int {
	var result int
	for _, value := range list1 {
		var times int
		for _, v := range list2 {
			if v == value {
				times = times + 1
			}
		}

		result = result + (value * times)
	}

	return result
}
