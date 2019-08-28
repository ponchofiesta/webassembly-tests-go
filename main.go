package main

import (
	"syscall/js"
	"unsafe"
	"webassembly_benchmarks_go/benchmarks"
	Hanoi "webassembly_benchmarks_go/benchmarks/hanoi/closure"
)

var (
	DATA_SORT_BASE  []benchmarks.User
	DATA_SORT       []benchmarks.User
	DATA_BYTES_BASE []byte
	DATA_BYTES      []byte
)

func main() {
	done := make(chan struct{})
	exports := make(map[string]interface{})

	exports["iterate"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//js.Global().Get("console").Call("debug", "Go: iterate")
		benchmarks.Iterate(args[0].Int())
		return nil
	})
	exports["strings_dynamic"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//js.Global().Get("console").Call("debug", "Go: strings_dynamic")
		return benchmarks.StringsDynamic(args[0].String(), args[1].String())
	})
	exports["strings_static"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//js.Global().Get("console").Call("debug", "Go: strings_dynamic")
		repeat := args[0].Int()
		a := "hello world"
		b := "world"
		for i := 0; i < repeat; i++ {
			benchmarks.StringsDynamic(a, b)
		}
		return nil
	})
	exports["sum"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//js.Global().Get("console").Call("debug", "Go: strings_dynamic")
		repeat := args[0].Int()
		for i := 0; i < repeat; i++ {
			benchmarks.Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
		return nil
	})
	exports["fibonacci"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: fobonacci")
		return benchmarks.Fibonacci(args[0].Int())
	})
	exports["hanoi"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: hanoi")
		hanoi := Hanoi.Hanoi()
		return hanoi(args[0].Int(), args[1].String(), args[2].String(), args[3].String())
	})
	exports["sort"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: sort")
		benchmarks.Sort(DATA_SORT)
		return nil
	})
	exports["prime"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: prime")
		benchmarks.Prime(uint64(args[0].Float()))
		return nil
	})
	exports["aes"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: aes")
		key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		iv := []byte{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
		benchmarks.AesEncrypt(key, iv, DATA_BYTES)
		return nil
	})
	exports["deflate"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: deflate")
		benchmarks.Deflate(DATA_BYTES)
		return nil
	})
	exports["convolve"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: convolve")
		matrix := []float64{
			0.0, 0.2, 0.0,
			0.2, 0.2, 0.2,
			0.0, 0.2, 0.0,
		}
		benchmarks.Convolve(args[0], matrix, 1)
		return nil
	})
	exports["convolve_mem"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: convolve_mem")
		ptr := uintptr(args[0].Int())
		width := args[1].Int()
		height := args[2].Int()
		size := width * height * 4
		data := (*[1 << 30]byte)(unsafe.Pointer(ptr))[:size:size]
		matrix := []float64{
			0.0, 0.2, 0.0,
			0.2, 0.2, 0.2,
			0.0, 0.2, 0.0,
		}
		benchmarks.ConvolveMem(data, width, height, matrix, 1)
		return nil
	})

	exports["prepare_test_data"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: prepare_test_data")
		switch args[0].String() {
		case "sort":
			var users []benchmarks.User
			jsUsers := args[1]
			len := jsUsers.Get("length").Int()
			for i := 0; i < len; i++ {
				user := benchmarks.User{
					Id:   jsUsers.Index(i).Get("id").Int(),
					Name: jsUsers.Index(i).Get("name").String(),
				}
				users = append(users, user)
			}
			DATA_SORT_BASE = users
		case "bytes":
			DATA_BYTES_BASE = []byte(args[1].String())
		}
		return nil
	})
	exports["reset_test_data"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: reset_test_data")
		switch args[0].String() {
		case "sort":
			DATA_SORT = make([]benchmarks.User, len(DATA_SORT_BASE))
			copy(DATA_SORT, DATA_SORT_BASE)
		case "bytes":
			DATA_BYTES = make([]byte, len(DATA_BYTES_BASE))
			copy(DATA_BYTES, DATA_BYTES_BASE)
		}
		return nil
	})
	exports["clear_test_data"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("debug", "Go: clear_test_data")
		switch args[0].String() {
		case "sort":
			DATA_SORT_BASE = nil
			DATA_SORT = nil
		case "aes":
			DATA_BYTES_BASE = nil
			DATA_BYTES = nil
		}
		return nil
	})

	js.Global().Get("wasm").Set("go", exports)

	<-done
}
