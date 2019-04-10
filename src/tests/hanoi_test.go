package tests

import "testing"
import "tests/hanoi/class_to_string"
import "tests/hanoi/class"
import "tests/hanoi/closure"

func BenchmarkHanoiBuilderClassToString(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		class_to_string.Hanoi(10, "A", "B", "C")
	}
	b.StopTimer()
}

func BenchmarkHanoiBuilderClass(b *testing.B) {
	hanoi := class.Hanoi{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hanoi.Hanoi(10, "A", "B", "C")
	}
	b.StopTimer()
}

func BenchmarkHanoiClosure(b *testing.B) {
	hanoi := closure.Hanoi()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hanoi(10, "A", "B", "C")
	}
	b.StopTimer()
}
