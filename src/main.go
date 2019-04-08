package main

import (
	"syscall/js"
	"tests"
)

func main() {
	done := make(chan struct{})
	exports := make(map[string]interface{})

	exports["fibonacci"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return tests.Fibonacci(args[0].Int())
	})
	exports["hanoi"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		hanoi := tests.Hanoi()
		return hanoi(args[0].Int(), args[1].String(), args[2].String(), args[3].String())
	})

	js.Global().Get("wasm").Set("go", exports)

	<-done
}
