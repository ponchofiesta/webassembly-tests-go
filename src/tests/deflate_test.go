package tests

import (
	"reflect"
	"testing"
)

func TestDeflate(t *testing.T) {

	text := "abcdefghijklmnop"
	data := []byte(text)

	expect := []byte{74, 76, 74, 78, 73, 77, 75, 207, 200, 204, 202, 206, 201, 205, 203, 47, 0, 4, 0, 0, 255, 255}
	compressed := Deflate(data)

	if !reflect.DeepEqual(expect, compressed) {
		t.Errorf("actual: %d, expect: %d", compressed, expect)
	}

	decompressed := Inflate(compressed)
	decompressedText := string(decompressed)
	if text != decompressedText {
		t.Errorf("actual: %s, expect: %s", decompressedText, text)
	}
}
