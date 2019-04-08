package tests

import "strings"

func Hanoi() (hanoi func(n int, from, via, to string) string) {
	var moves strings.Builder
	hanoi = func(n int, from, via, to string) string {
		if n > 0 {
			hanoi(n-1, from, to, via)
			moves.WriteString(from + "->" + to + "\n")
			hanoi(n-1, via, from, to)
		}
		return moves.String()
	}
	return
}
