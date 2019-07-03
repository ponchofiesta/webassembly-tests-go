package benchmarks

import "strings"

func StringsDynamic(a string, b string) bool {
	return strings.Contains(a, b)
}
