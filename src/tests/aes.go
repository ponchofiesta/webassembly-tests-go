package tests

import (
	"crypto/aes"
	"crypto/cipher"
)

func Aes(key []uint8, iv []uint8, data []uint8) []uint8 {
	c, _ := aes.NewCipher(key)
	encrypter := cipher.NewCBCEncrypter(c, iv)
	encrypted := make([]byte, len(data))
	encrypter.CryptBlocks(encrypted, data)
	return encrypted
}
