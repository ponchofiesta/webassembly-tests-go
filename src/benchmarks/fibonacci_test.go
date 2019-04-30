package benchmarks

import "testing"

func TestFibonacci(t *testing.T) {

	expect := 6765
	actual := Fibonacci(20)

	if actual != 6765 {
		t.Errorf("actual: %d, expect: %d", actual, expect)
	}
}
