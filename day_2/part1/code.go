package part1

import (
	"math"
)

func Solve(report []int) bool {
	prev := 0
	increase := false

	if report[0] < report[1] {
		increase = true
	}

	for _, value := range report {
		if prev == 0 {
			prev = value
			continue
		}

		// Must not be the same value
		if prev == value {
			return false
		}

		// Any two adjacent levels differ by at least one and at most three
		if distance(prev, value) > 3 {
			return false
		}

		// The levels are either all increasing or all decreasing.
		if (increase && prev > value) || (!increase && prev < value) {
			return false
		}

		prev = value
	}
	return true
}

func distance(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
