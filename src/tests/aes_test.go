package tests

import (
	"reflect"
	"testing"
)

func TestAes(t *testing.T) {

	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	iv := []byte{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	data := []byte("abcdefghijklmnop")

	expect := []byte{}
	actual := Aes(key, iv, data)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("actual: %d, expect: %d", actual, expect)
	}
}
