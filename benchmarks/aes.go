package benchmarks

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesEncrypt(key []uint8, iv []uint8, data []uint8) []uint8 {
	c, _ := aes.NewCipher(key)
	encrypter := cipher.NewCBCEncrypter(c, iv)
	encrypted := make([]byte, len(data))
	encrypter.CryptBlocks(encrypted, data)
	return encrypted
}

func AesDecrypt(key []uint8, iv []uint8, data []uint8) []uint8 {
	c, _ := aes.NewCipher(key)
	encrypter := cipher.NewCBCDecrypter(c, iv)
	decrypted := make([]byte, len(data))
	encrypter.CryptBlocks(decrypted, data)
	return decrypted
}
