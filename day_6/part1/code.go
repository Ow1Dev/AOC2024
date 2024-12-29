package part1

import (
	"log"
	"strings"
)

const obstacle = '#'
const wall = 'P'
const marker = 'X'

func Solve(input string) int {
	width := getWidth(input) + 2
	n := strings.Repeat(string(wall), width)
	for _, line := range strings.Split(input, "\n") {
		n += string(wall) + line + string(wall)
	}
	n += strings.Repeat(string(wall), width)
	m := []rune(n)

	cursor := strings.Index(n, "^")
	m[cursor] = marker

	d := 0

	for {
		nextcursor := getNextCursor(d, cursor, width)
		if m[nextcursor] == wall {
			break
		}
		if m[nextcursor] == obstacle {
			d = (d + 1) % 4
			nextcursor = getNextCursor(d, cursor, width)
		}

		m[nextcursor] = marker
		cursor = nextcursor
	}

	result := strings.Count(string(m), string(marker))

	return result
}

func getNextCursor(d, cursor, width int) int {
	switch d {
	case 0:
		return cursor - width
	case 1:
		return cursor + 1
	case 2:
		return cursor + width
	case 3:
		return cursor - 1
	default:
		log.Panicln("Direction is not there")
	}

	return -1
}

func getWidth(input string) int {
	for i := range input {
		if input[i] == '\n' {
			return i
		}
	}

	return len(input)
}
