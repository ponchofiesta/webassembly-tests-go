package tests

import "testing"

func TestHanoi(t *testing.T) {

	expect := "A->C\nA->B\nC->B\nA->C\nB->A\nB->C\nA->C\n"
	hanoi := Hanoi()
	actual := hanoi(3, "A", "B", "C")

	if actual != expect {
		t.Errorf("actual: %s, expected: %s", actual, expect)
	}
}
