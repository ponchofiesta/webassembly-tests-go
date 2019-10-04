package benchmarks

import (
	"reflect"
	"testing"
)

func TestSha256(t *testing.T) {

	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	expect := "66840dda154e8a113c31dd0ad32f7f3a366a80e8136979d8f5a101d3d29d6f72"
	actual := Sha256(data)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("actual: %d, expect: %d", actual, expect)
	}
}
