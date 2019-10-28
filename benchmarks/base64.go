package benchmarks

import "encoding/base64"

func Base64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
