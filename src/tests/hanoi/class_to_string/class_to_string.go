package class_to_string

type move struct {
	from string
	to   string
}

func (m move) String() string {
	return m.from + "->" + m.to + "\n"
}

type moves []move

func (m moves) String() (str string) {
	for _, s := range m {
		str += s.String()
	}
	return
}

func hanoi(moves moves, n int, from, via, to string) moves {
	if n > 0 {
		moves = hanoi(moves, n-1, from, to, via)
		moves = append(moves, move{from, to})
		moves = hanoi(moves, n-1, via, from, to)
	}
	return moves
}

func Hanoi(n int, from, via, to string) string {
	var moves moves
	return hanoi(moves, n, from, via, to).String()
}
