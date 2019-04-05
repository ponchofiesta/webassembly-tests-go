package main

import "testing"

func TestFibonacci(t *testing.T) {
	got := fibonacci(20)
	if got != 6765 {
		t.Errorf("fibonacci=%d, want 6765", got)
	}
}
