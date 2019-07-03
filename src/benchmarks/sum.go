package benchmarks

func Sum(numbers []int) int {
	sum := 0
	for _, x := range numbers {
		sum += x
	}
	return sum
}
