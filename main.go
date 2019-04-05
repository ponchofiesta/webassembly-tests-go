package main

import (
	"syscall/js"
)

func main() {
	done := make(chan struct{})
	exports := make(map[string]interface{})

	exports["fibonacci"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fibonacci(args[0].Int())
	})

	js.Global().Get("wasm").Set("go", exports)

	<-done
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	a := 0
	b := 1
	for n--; n > 0; n-- {
		a = a + b
		a, b = b, a
	}
	return b
}
