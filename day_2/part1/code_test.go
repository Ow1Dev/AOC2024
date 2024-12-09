package part1

import "testing"

func TestSolveSolution(t *testing.T) {
  	tests := []struct {
		name     string
		report   []int
		expected bool
	}{
		{"safe", []int{7, 6, 4, 2, 1}, true},
		{"unsafe", []int{1, 2, 7, 8, 9}, false},
		{"unsafe", []int{9, 7, 6, 2, 1}, false},
		{"unsafe", []int{1, 3, 2, 4, 5}, false},
		{"unsafe", []int{8, 6, 4, 4, 1}, false},
		{"safe", []int{1, 3, 6, 7, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Solve(tt.report)
			if actual != tt.expected {
				t.Errorf("%s: want %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
