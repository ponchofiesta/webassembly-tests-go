package benchmarks

import (
	"reflect"
	"testing"
)

func TestAes(t *testing.T) {

	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	iv := []byte{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	text := "abcdefghijklmnop"
	data := []byte(text)

	expect := []byte{239, 202, 78, 31, 15, 43, 252, 66, 190, 102, 119, 142, 185, 132, 167, 184}
	encrypted := AesEncrypt(key, iv, data)

	if !reflect.DeepEqual(expect, encrypted) {
		t.Errorf("actual: %d, expect: %d", encrypted, expect)
	}

	decrypted := AesDecrypt(key, iv, encrypted)
	decryptedText := string(decrypted)
	if text != decryptedText {
		t.Errorf("actual: %s, expect: %s", decryptedText, expect)
	}
}
