package benchmarks

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {

	expect := []uint64{2, 3, 5, 7, 11, 13, 17, 19}
	actual := Prime(20)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("actual: %d, expect: %d", actual, expect)
	}
}
