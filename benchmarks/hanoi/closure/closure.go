package closure

import "strings"

func Hanoi() (hanoi func(n int, from, via, to string) string) {
	var moves strings.Builder
	var hanoiRe func(n int, from, via, to string)
	hanoiRe = func(n int, from, via, to string) {
		if n > 0 {
			hanoiRe(n-1, from, to, via)
			moves.WriteString(from)
			moves.WriteString("->")
			moves.WriteString(to)
			moves.WriteString("\n")
			hanoiRe(n-1, via, from, to)
		}
	}
	hanoi = func(n int, from, via, to string) string {
		hanoiRe(n, from, via, to)
		return moves.String()
	}
	return
}
