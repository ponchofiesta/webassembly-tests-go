package main

import (
	"syscall/js"
	"webassembly_benchmarks_go/benchmarks"
)

var (
	DATA_SORT_BASE  []benchmarks.User
	DATA_SORT       []benchmarks.User
	DATA_BYTES_BASE []byte
	DATA_BYTES      []byte
)

func copyByteArrayToGo(input js.Value) []byte {
	uint8arrayWrapper := js.Global().Get("Uint8Array").New(input)
	data := make([]byte, uint8arrayWrapper.Get("byteLength").Int())
	reader := js.TypedArrayOf(data)
	reader.Call("set", uint8arrayWrapper)
	reader.Release()
	return data
}

func copyByteArrayToJs(input []byte) js.Value {
	writer := js.TypedArrayOf(input)
	data := js.Global().Get("Uint8Array").New(len(input))
	data.Call("set", writer)
	writer.Release()
	return data
}

func copyUin32ArrayToGo(input js.Value) []uint32 {
	uint32arrayWrapper := js.Global().Get("Uint32Array").New(input)
	data := make([]uint32, uint32arrayWrapper.Get("byteLength").Int())
	reader := js.TypedArrayOf(data)
	reader.Call("set", uint32arrayWrapper)
	reader.Release()
	return data
}

func copyUint32ArrayToJs(input []uint32) js.Value {
	writer := js.TypedArrayOf(input)
	data := js.Global().Get("Uint32Array").New(len(input))
	data.Call("set", writer)
	writer.Release()
	return data
}
