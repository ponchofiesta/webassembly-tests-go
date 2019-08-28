package class

import "strings"

type Hanoi struct {
	moves strings.Builder
}

func (hanoi *Hanoi) hanoiRe(n int, from, via, to string) {
	if n > 0 {
		hanoi.hanoiRe(n-1, from, to, via)
		hanoi.moves.WriteString(from)
		hanoi.moves.WriteString("->")
		hanoi.moves.WriteString(to)
		hanoi.moves.WriteString("\n")
		hanoi.hanoiRe(n-1, via, from, to)
	}
}

func (hanoi *Hanoi) Hanoi(n int, from, via, to string) string {
	hanoi.hanoiRe(n, from, via, to)
	return hanoi.moves.String()
}
