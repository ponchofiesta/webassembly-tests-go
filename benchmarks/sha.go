package benchmarks

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Sha256(data []byte) string {
	result := sha256.Sum256(data)
	return fmt.Sprintf("%x\n", result)
}

func Sha512(data []byte) string {
	result := sha512.Sum512(data)
	return fmt.Sprintf("%x\n", result)
}
