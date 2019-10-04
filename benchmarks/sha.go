package benchmarks

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(data []byte) string {
	result := sha256.Sum256(data)
	return fmt.Sprintf("%x\n", result)
}
