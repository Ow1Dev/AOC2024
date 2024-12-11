package part1

import (
	"strconv"
	"unicode"
)

type mulNode struct {
	A int
	B int
}

func Solve(code string) int {
	nodes := parse(code)
	var result int

	for _, node := range nodes {
		result += node.A * node.B
	}

	return result
}

func parse(code string) []mulNode {
	nodes := make([]mulNode, 0)

	for {
		if len(code) >= 4 && code[0:4] == "mul(" {
			node := parseMulNode(code[4:])
			if node != nil {
				nodes = append(nodes, *node)
			}
		}

		code = code[1:]
		if len(code) == 0 {
			break
		}
	}

	return nodes
}

func parseMulNode(code string) *mulNode {
	var aStr string
	var bStr string

	var switcharg = false
	var isEnding = false

	for {
		if unicode.IsDigit(rune(code[0])) {
			if switcharg == false {
				aStr += string(code[0])
			} else {
				bStr += string(code[0])
			}
		} else if code[0] == ')' {
			isEnding = true
			break
		} else if code[0] == ',' && !switcharg {
			switcharg = true
		} else {
			return nil
		}

		code = code[1:]
		if len(code) == 0 {
			return nil
		}
	}

	if !isEnding {
		return nil
	}

	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	return &mulNode{
		A: a,
		B: b,
	}
}
