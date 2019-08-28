package class_to_string

import "testing"

func TestHanoi(t *testing.T) {
	expect := "A->C\nA->B\nC->B\nA->C\nB->A\nB->C\nA->C\n"
	actual := Hanoi(3, "A", "B", "C")

	if actual != expect {
		t.Errorf("actual: %s, expected: %s", actual, expect)
	}
}
