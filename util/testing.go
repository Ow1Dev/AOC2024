package util

import "testing"

func AssertEqual(t *testing.T, got, want int, message string) {
	if got != want {
		t.Errorf("%s: got %d, expected %d", message, got, want)
	}
}
