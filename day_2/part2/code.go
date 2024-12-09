package part2

import (
	"math"
	"slices"
)

func Solve(report []int) bool {
	if isReportSafe(report) {
		return true
	}

  for i := 0; i < len(report); i++ {
    r := slices.Clone(report)
    r = append(r[:i], r[i+1:]...)

    if isReportSafe(r) {
      return true
    }
  }

	return false
}

func isReportSafe(report []int) bool {
	increase := false
	if report[0] < report[1] {
		increase = true
	}

	prev := report[0]
	for index, value := range report {
		if index == 0 {
			continue
		}

		if isSafe(prev, value, increase) == false {
			return false
		}

		prev = value
	}

	return true
}

func isSafe(prev, curr int, increase bool) bool {
	// Must not be the same value
	if prev == curr {
		return false
	}

	// Any two adjacent levels differ by at least one and at most three
	if distance(prev, curr) > 3 {
		return false
	}

	// The levels are either all increasing or all decreasing.
	if (increase && prev > curr) || (!increase && prev < curr) {
		return false
	}

	return true
}

func distance(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func removeAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
