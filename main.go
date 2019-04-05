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
	exports["hanoi"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		hanoi(args[0].Int(), args[1].Int(), args[2].Int(), args[3].Int())
		return nil
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
		a += b
		a, b = b, a
	}
	return b
}

func hanoi(n int, from, to, via int) {
	if n > 0 {
		hanoi(n-1, from, via, to)
		hanoi(n-1, to, from, via)
	}
}
